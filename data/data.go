package data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

func NewRandomData() Data {
	return Data{
		Bool:    gofakeit.Bool(),
		Int64:   gofakeit.Int64(),
		String:  utils.RandomString(),
		Float64: gofakeit.Float64(),
	}
}

type Data struct {
	Bool    bool
	Int64   int64
	String  string
	Float64 float64
}
