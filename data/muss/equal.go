package mus

func EqualData(data1, data2 Data) bool {
	return data1.Bool == data2.Bool && data1.Int64 == data2.Int64 &&
		data1.String == data2.String &&
		data1.Float64 == data2.Float64
}
