package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Database struct {
	Tables []string
	Data   map[string][]map[string]interface{}
}

func DatabaseFactory(path string) Database {
	db := Database{}
	db.JsonToDatabase(
		readJsonFile(path),
	)

	for key := range db.Data {
		db.Tables = append(db.Tables, key)
	}

	return db
}

func (db *Database) JsonToDatabase(byteValue []byte) {
	json.Unmarshal([]byte(byteValue), &db.Data)
}

func readJsonFile(path string) []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
