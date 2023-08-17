package kitex

import "github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"

func EqualData(data1, data2 *echo.KitexData) bool {
	return data1.Bool == data2.Bool && data1.Int64 == data2.Int64 &&
		data1.String_ == data2.String_ &&
		data1.Float64 == data2.Float64
}
