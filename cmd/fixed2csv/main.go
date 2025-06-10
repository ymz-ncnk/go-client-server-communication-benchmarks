package main

import (
	"os"
	"path/filepath"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd"
)

const (
	SecFileName   = "sec.csv"
	CopFileName   = "cop.csv"
	MaxFileName   = "max.csv"
	MedFileName   = "med.csv"
	MinFileName   = "min.csv"
	P99FileName   = "p99.csv"
	BFileName     = "b.csv"
	AllocFileName = "allocs.csv"
)

func main() {
	inputFile, outputDir := cmd.ParseFlags()

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parser := NewParser(file)
	nsData, copData, maxData, medData, minData, p99Data, bData, allocData, err :=
		parser.Parse()
	if err != nil {
		panic(err)
	}
	var (
		metricData = BenchToMetricData(nsData, float64(1e9))
		csvData    = MetricToCSVData(metricData)
		path       = filepath.Join(outputDir, SecFileName)
	)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(copData, float64(1e9))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, CopFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(maxData, float64(1e9))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, MaxFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(medData, float64(1e9))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, MedFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(minData, float64(1e9))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, MinFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(p99Data, float64(1e9))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, P99FileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(bData, float64(1e3))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, BFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}

	metricData = BenchToMetricData(allocData, float64(1))
	csvData = MetricToCSVData(metricData)
	path = filepath.Join(outputDir, AllocFileName)
	if err = cmd.WriteCSV(path, csvData); err != nil {
		panic(err)
	}
}
