package api

import (
	"fmt"
	"json-server/database"
	"net"
	"strconv"
)

type Api struct {
	db        *database.Database
	endpoints []Endpoint
}

func ApiFactory(jsonfile string) Api {

	db := database.DatabaseFactory(jsonfile)
	var endpoints []Endpoint
	for _, value := range db.Tables {
		endpoints = append(endpoints, EndpointFactory(value, &db))
	}
	return Api{&db, endpoints}
}

func ValidatePort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return fmt.Errorf("invalid port")
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("port in use")
	}

	_ = ln.Close()

	return nil
}
