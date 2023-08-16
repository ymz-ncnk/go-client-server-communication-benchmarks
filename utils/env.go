package utils

import (
	"os"
	"strconv"
)

const DefClientsCount = 1
const DefGenSize = 200000

func ClientsCount() int {
	val := os.Getenv("CLIENTS_COUNT")
	if val == "" {
		return DefClientsCount
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	if n == 0 {
		panic("CLIENTS_COUNT == 0")
	}
	return n
}

func GenSize() int {
	val := os.Getenv("GEN_SIZE")
	if val == "" {
		return DefGenSize
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return n
}
