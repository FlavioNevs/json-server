package api

import (
	"fmt"
	"json-server/database"
	"net"
	"strconv"
)

type Api struct {
	db *database.Database
}

func ApiFactory(db *database.Database) Api {
	return Api{db}
}

func ValidPort(port string) error {
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
