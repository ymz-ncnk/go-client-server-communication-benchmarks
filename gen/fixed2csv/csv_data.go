package main

import "strconv"

type CSVData struct {
	headers []string
	values  [][]float64
}

func (d CSVData) Headers() []string {
	return d.headers
}

func (d CSVData) Values() [][]float64 {
	return d.values
}

func (d CSVData) ValueToString(val float64, col int) string {
	if col == 0 {
		return strconv.Itoa(int(val))
	}
	
	return strconv.FormatFloat(val, 'f', 3, 64)
}
