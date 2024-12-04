package kitex

import (
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/data"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
)

type DataSet = [][]*echo.KitexData

func ConvertDataSet(ds data.DataSet) (ads DataSet) {
	ads = make([][]*echo.KitexData, len(ds))
	for i := 0; i < len(ds); i++ {
		ads[i] = make([]*echo.KitexData, len(ds[i]))
		for j := 0; j < len(ds[i]); j++ {
			ads[i][j] = &echo.KitexData{
				Bool:    ds[i][j].Bool,
				Int64:   ds[i][j].Int64,
				String_: ds[i][j].String,
				Float64: ds[i][j].Float64,
			}
		}
	}
	return
}

func EqualData(d1, d2 *echo.KitexData) bool {
	return d1.Bool == d2.Bool && d1.Int64 == d2.Int64 &&
		d1.String_ == d2.String_ &&
		d1.Float64 == d2.Float64
}
