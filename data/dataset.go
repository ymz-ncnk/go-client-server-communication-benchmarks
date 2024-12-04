package data

import (
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/mus"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
)

func GenerateDataSet(clientsCount, size int) (s DataSet) {
	s = make([][]Data, clientsCount)
	for i := 0; i < clientsCount; i++ {
		s[i] = make([]Data, size)
		for j := 0; j < size; j++ {
			s[i][j] = NewRandomData()
		}
	}
	return
}

type DataSet [][]Data

func (s DataSet) ToMUS() (as mus.DataSetMUS) {
	as = make([][]mus.Data, len(s))
	for i := 0; i < len(s); i++ {
		as[i] = make([]mus.Data, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			as[i][j] = mus.Data{
				Bool:    s[i][j].Bool,
				Int64:   s[i][j].Int64,
				String:  s[i][j].String,
				Float64: s[i][j].Float64,
			}
		}
	}
	return
}

func (s DataSet) ToProtobuf() (as protobuf.DataSetProtobuf) {
	as = make([][]*protobuf.Data, len(s))
	for i := 0; i < len(s); i++ {
		as[i] = make([]*protobuf.Data, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			as[i][j] = &protobuf.Data{
				Bool:    s[i][j].Bool,
				Int64:   s[i][j].Int64,
				String_: s[i][j].String,
				Float64: s[i][j].Float64,
			}
		}
	}
	return
}
