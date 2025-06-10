package main

import (
	"sort"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"
)

type QPS float64

type QPSData map[BenchCase]QPS

func QPSToCSVData(qpsData QPSData) CSVData {
	var (
		headers = map[string]struct{}{}
		maxN    int
	)
	for c := range qpsData {
		headers[c.Name] = struct{}{}
		if c.N > maxN {
			maxN = c.N
		}
	}
	data := CSVData{
		headers: make([]string, 0, len(headers)+1),
		values:  make([][]int, gen.Log2Convert(maxN)+1),
	}
	data.headers = append(data.headers, "clients")
	for header := range headers {
		data.headers = append(data.headers, header)
	}
	sort.Strings(data.headers[1:])
	for i := range data.values {
		data.values[i] = make([]int, len(data.headers))
	}

	for c, qps := range qpsData {
		var (
			i      int
			header string
		)
		for i, header = range data.headers {
			if c.Name == header {
				break
			}
		}
		clients := gen.Log2Convert(c.N)
		data.values[clients][0] = c.N
		data.values[clients][i] = int(qps)
	}
	return data
}
