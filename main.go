package main

import (
	"fmt"
	"json-server/api"
	"os"
)

var (
	BASE_DIR string
	Port     string
	JsonFile string
)

func init() {
	// Setting Current Workdir
	base_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BASE_DIR = base_dir

	fmt.Print("Json File: ")
	fmt.Scan(&JsonFile)
	if err := FindFile(&JsonFile); err != nil {
		panic(err)
	}

	fmt.Print("Port: ")
	fmt.Scan(&Port)
	if err = api.ValidPort(Port); err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println(JsonFile)
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
