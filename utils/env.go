package utils

import (
	"os"
	"strconv"
)

func ClientsCount() (clientsCount int) {
	val := os.Getenv("CLIENTS_COUNT")
	if val == "" {
		return 1
	}
	clientsCount, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return
}
