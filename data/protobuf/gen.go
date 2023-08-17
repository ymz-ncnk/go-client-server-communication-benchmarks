package protobuf

import (
	"math/rand"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

func GenData(clientsCount, size int, r *rand.Rand) (arr [][]*Data) {
	arr = make([][]*Data, clientsCount)
	for i := 0; i < len(arr); i++ {
		sub := make([]*Data, size)
		for j := 0; j < len(sub); j++ {
			sub[j] = NewRandomData(r)
		}
		arr[i] = sub
	}
	return
}

func NewRandomData(r *rand.Rand) *Data {
	return &Data{
		Bool:    utils.RandomBool(r),
		Int64:   utils.RandomInt64(r),
		String_: utils.RandomString(r),
		Float64: utils.RandomFloat64(r),
	}
}
