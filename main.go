package main

import (
	"fmt"
	"json-server/api"
	"json-server/database"
	"net/http"
	"os"
	"time"
)

var (
	BASE_DIR string
	Port     string
	JsonFile string
	Api      api.Api
	Db       database.Database
)

func init() {
	// Setting Current Workdir
	base_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BASE_DIR = base_dir

	// Setting JsonFile Path
	fmt.Print("Json File: ")
	fmt.Scan(&JsonFile)
	if err := FindFile(&JsonFile); err != nil {
		panic(err)
	}

	// Setting API Port
	fmt.Print("Port: ")
	fmt.Scan(&Port)
	if err = api.ValidatePort(Port); err != nil {
		panic(err)
	}
}

func main() {
	PrintLabel()
	Api = api.ApiFactory(JsonFile)
	http.ListenAndServe(":"+Port, Api.RouterFactory())
}

func PrintLabel() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(time.Now().Format("January 02, 2006 - 15:04:05"))
	fmt.Println("Creating endpoints from " + JsonFile)
	fmt.Println("Starting development server at http://127.0.0.1:" + Port + "/")
	fmt.Println("Quit the server with CTRL-BREAK.")
	fmt.Println("")
}

func FindFile(filepath *string) error {

	if _, err := os.Stat(*filepath); !os.IsNotExist(err) {
		return nil
	}

	new_filepath := fmt.Sprintf("%s/%s", BASE_DIR, *filepath)
	if _, err := os.Stat(*filepath); !os.IsNotExist(err) {
		filepath = &new_filepath
		return nil
	}

	return fmt.Errorf("json file not found")
}
