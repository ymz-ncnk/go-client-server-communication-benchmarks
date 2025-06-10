package main

import (
	"os"
	"path/filepath"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"
)

const (
	FileName = "qps.csv"
)

func main() {
	inputFile, outputDir := gen.ParseFlags()

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parser := NewParser(file)
	benchData, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	var (
		qpsData = BenchToQPSData(benchData)
		csvData = QPSToCSVData(qpsData)
		path    = filepath.Join(outputDir, FileName)
	)
	if err = gen.WriteCSV(path, csvData); err != nil {
		panic(err)
	}
}
