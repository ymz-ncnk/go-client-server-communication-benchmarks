package main

import "github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"

type BenchData map[gen.BenchCase][]int

func BenchToMetricData(data BenchData, div float64) MetricData {
	metrics := make(MetricData)
	for benchCase, values := range data {
		if len(values) == 0 {
			continue // avoid division by zero
		}
		var sum int
		for _, v := range values {
			sum += v
		}
		var avg = (float64(sum) / div) / float64(len(values))
		metrics[benchCase] = avg
	}
	return metrics
}
