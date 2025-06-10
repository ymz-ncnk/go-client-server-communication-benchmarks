package main

import "strconv"

type CSVData struct {
	headers []string
	values  [][]int
}

func (d CSVData) Headers() []string {
	return d.headers
}

func (d CSVData) Values() [][]int {
	return d.values
}

func (d CSVData) ValueToString(val int, col int) string {
	return strconv.Itoa(val)
}
