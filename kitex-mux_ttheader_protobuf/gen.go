package kitex

import (
	"math/rand"

	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
)

func GenData(clientsCount, size int, r *rand.Rand) (arr [][]*echo.KitexData) {
	arr = make([][]*echo.KitexData, clientsCount)
	for i := 0; i < len(arr); i++ {
		sub := make([]*echo.KitexData, size)
		for j := 0; j < len(sub); j++ {
			sub[j] = NewRandomData(r)
		}
		arr[i] = sub
	}
	return
}

func NewRandomData(r *rand.Rand) *echo.KitexData {
	return &echo.KitexData{
		Bool:    utils.RandomBool(r),
		Int64:   utils.RandomInt64(r),
		String_: utils.RandomString(r),
		Float64: utils.RandomFloat64(r),
	}
}
