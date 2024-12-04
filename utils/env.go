package utils

import (
	"os"
	"strconv"
)

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
