package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"
)

var lineRe = regexp.MustCompile(
	`^BenchmarkFixed/(\d+)/([\w.\-]+)-\d+\s+` + // Name + N
		`\d+\s+` + // Iters (unused)
		`\d+ N\s+` + // N (unused)
		`(\d+) ns\s+` + // Ns
		`(\d+) ns/cop\s+` + // Cop
		`(\d+) ns/max\s+` + // Max
		`(\d+) ns/med\s+` + // Med
		`(\d+) ns/min\s+` + // Min
		`(\d+) ns/p99\s+` + // P99
		`(\d+) B/op\s+` + // B/op
		`(\d+) allocs/op`, // Allocs
)

func NewParser(file *os.File) Parser {
	return Parser{
		file:      file,
		nsData:    make(BenchData),
		copData:   make(BenchData),
		maxData:   make(BenchData),
		medData:   make(BenchData),
		minData:   make(BenchData),
		p99Data:   make(BenchData),
		bData:     make(BenchData),
		allocData: make(BenchData),
	}
}

type Parser struct {
	file      *os.File
	nsData    BenchData
	copData   BenchData
	maxData   BenchData
	medData   BenchData
	minData   BenchData
	p99Data   BenchData
	bData     BenchData
	allocData BenchData
}

func (p Parser) Parse() (nsData, copData, maxData, medData, minData, p99Data,
	bData, allocData BenchData, err error) {
	var (
		scanner = bufio.NewScanner(p.file)
		ok      bool
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.HasPrefix(line, "Benchmark") {
			continue
		}
		if err = p.parseLine(line); err != nil {
			return
		}
		ok = true
	}
	if err = scanner.Err(); err != nil {
		return
	}
	if !ok {
		err = gen.NewNotBenchmarksFileError(p.file.Name())
		return
	}
	return p.nsData, p.copData, p.maxData, p.medData, p.minData, p.p99Data, p.bData, p.allocData, nil
}

func (p Parser) parseLine(line string) error {
	matches := lineRe.FindStringSubmatch(line)
	if matches == nil {
		return fmt.Errorf("no match: %s", line)
	}
	name := matches[2]
	n, _ := strconv.Atoi(matches[1])
	key := gen.BenchCase{Name: name, N: n}

	// Parse values
	ns, _ := strconv.Atoi(matches[3])
	copNs, _ := strconv.Atoi(matches[4])
	maxNs, _ := strconv.Atoi(matches[5])
	medNs, _ := strconv.Atoi(matches[6])
	minNs, _ := strconv.Atoi(matches[7])
	p99Ns, _ := strconv.Atoi(matches[8])
	bOp, _ := strconv.Atoi(matches[9])
	allocsOp, _ := strconv.Atoi(matches[10])

	// Append to slices
	p.nsData[key] = append(p.nsData[key], ns)
	p.copData[key] = append(p.copData[key], copNs)
	p.maxData[key] = append(p.maxData[key], maxNs)
	p.medData[key] = append(p.medData[key], medNs)
	p.minData[key] = append(p.minData[key], minNs)
	p.p99Data[key] = append(p.p99Data[key], p99Ns)
	p.bData[key] = append(p.bData[key], bOp)
	p.allocData[key] = append(p.allocData[key], allocsOp)
	return nil
}
