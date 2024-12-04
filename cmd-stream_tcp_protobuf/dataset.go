package csp

import (
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/data"
)

type DataSet [][]EchoCmd

func ConvertDataSet(ds data.DataSet) (ads DataSet) {
	pds := ds.ToProtobuf()
	ads = make([][]EchoCmd, len(ds))
	for i := 0; i < len(ds); i++ {
		ads[i] = make([]EchoCmd, len(ds[i]))
		for j := 0; j < len(ds[i]); j++ {
			ads[i][j] = EchoCmd{pds[i][j]}
		}
	}
	return
}
