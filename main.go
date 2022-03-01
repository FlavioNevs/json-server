package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ()

var (
	config   Api
	database map[string][]map[string]interface{}
	result   Base
)

func main() {
	load_json("D:/Dev/GO/json-server/base.json")
	log.Fatal(http.ListenAndServe(":8000", config.CreateRouter()))
}

func load_json(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result)

	database = result.Data
	config = result.Api
}
