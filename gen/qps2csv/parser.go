package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/gen"
)

var lineRe = regexp.MustCompile(`^BenchmarkQPS/(\d+)/([^\s]+)-\d+\s+(\d+)\s+(\d+)\s+ns$`)

func NewParser(file *os.File) Parser {
	return Parser{file: file, benchData: make(BenchData)}
}

type Parser struct {
	file      *os.File
	benchData BenchData
}

func (p Parser) Parse() (benchData BenchData, err error) {
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
	return p.benchData, nil
}

func (p Parser) parseLine(line string) (err error) {
	matches := lineRe.FindStringSubmatch(line)
	if matches == nil {
		return fmt.Errorf("no match: %s", line)
	}
	var (
		name = matches[2]
		key  = BenchCase{Name: name, N: gen.StrToInt(matches[1])}
	)
	p.benchData[key] = append(p.benchData[key], BenchEntry{
		Iter: gen.StrToInt(matches[3]),
		Ns:   gen.StrToInt64(matches[4]),
	})
	return
}
