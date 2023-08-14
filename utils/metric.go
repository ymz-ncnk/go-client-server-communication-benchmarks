package utils

import (
	"testing"
	"time"

	"github.com/montanaflynn/stats"
)

func ReportMetrics(clientsCount int, copsD chan time.Duration, b *testing.B) {
	copsDArr := makeCopsDArr(copsD)

	mean, _ := stats.Mean(copsDArr)
	med, _ := stats.Median(copsDArr)
	max, _ := stats.Max(copsDArr)
	min, _ := stats.Min(copsDArr)
	p99, _ := stats.Percentile(copsDArr, 99.9)

	b.ReportMetric(0, "ns/op")
	b.ReportMetric(float64(b.Elapsed()), "ns")
	b.ReportMetric(mean, "ns/cop")
	b.ReportMetric(med, "ns/med")
	b.ReportMetric(max, "ns/max")
	b.ReportMetric(min, "ns/min")
	b.ReportMetric(p99, "ns/p99")
}

func makeCopsDArr(copsD chan time.Duration) (copsDArr []float64) {
	close(copsD)
	copsDArr = []float64{}
	for spent := range copsD {
		copsDArr = append(copsDArr, float64(spent))
	}
	return
}
