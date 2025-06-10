package main

import (
	"sort"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"
)

type MetricData map[gen.BenchCase]float64

func MetricToCSVData(data MetricData) CSVData {
	var (
		headers = map[string]struct{}{}
		maxN    int
	)
	for c := range data {
		headers[c.Name] = struct{}{}
		if c.N > maxN {
			maxN = c.N
		}
	}
	csv := CSVData{
		headers: make([]string, 0, len(headers)+1),
		values:  make([][]float64, gen.Log2Convert(maxN)+1),
	}
	csv.headers = append(csv.headers, "clients")
	for header := range headers {
		csv.headers = append(csv.headers, header)
	}
	sort.Strings(csv.headers[1:]) // sort benchmark names, skip "clients"
	for i := range csv.values {
		csv.values[i] = make([]float64, len(csv.headers))
	}
	for c, value := range data {
		var colIndex int
		for i, header := range csv.headers {
			if c.Name == header {
				colIndex = i
				break
			}
		}
		rowIndex := gen.Log2Convert(c.N)
		csv.values[rowIndex][0] = float64(c.N)
		csv.values[rowIndex][colIndex] = value
	}
	return csv
}
