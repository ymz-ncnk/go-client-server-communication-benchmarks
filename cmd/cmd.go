package cmd

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

type BenchCase struct {
	Name string
	N    int
}

type CSVData[T any] interface {
	Headers() []string
	Values() [][]T
	ValueToString(val T, col int) string
}

func ParseFlags() (inputFile, outputDir string) {
	input := flag.String("i", "", "Path to input file")
	output := flag.String("d", "", "Path to output directory")
	flag.Parse()

	inputFile = *input
	outputDir = *output

	if inputFile == "" || outputDir == "" {
		fmt.Println("Usage: program -i input.txt -d output/")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if output directory exists
	if stat, err := os.Stat(outputDir); err != nil || !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %q is not a valid directory\n", outputDir)
		os.Exit(1)
	}
	return
}

func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func StrToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func Log2Convert(n int) int {
	if n <= 0 || (n&(n-1)) != 0 {
		panic("input must be a power of 2 and greater than 0")
	}
	return int(math.Log2(float64(n)))
}

func WriteCSV[T any](filename string, data CSVData[T]) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(data.Headers()); err != nil {
		return fmt.Errorf("error writing headers: %w", err)
	}
	for _, row := range data.Values() {
		stringRow := make([]string, len(row))
		for i, val := range row {
			stringRow[i] = data.ValueToString(val, i)
		}
		if err := writer.Write(stringRow); err != nil {
			return fmt.Errorf("error writing row: %w", err)
		}
	}
	return nil
}

func NewNotBenchmarksFileError(fileName string) error {
	return fmt.Errorf("there no benchmarks data in %v", fileName)
}
