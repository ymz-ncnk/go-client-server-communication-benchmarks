#!/bin/bash
GEN_SIZE=400000 go test -bench BenchmarkQPS -count=10 -timeout=30m > ./results/qps/benchmarks.txt
go test -bench BenchmarkFixed -benchtime=100000x -benchmem -count=10 > ./results/fixed/benchmarks.txt

if [ ! -e "fixed2csv" ]; then
	go build -o fixed2csv ./cmd/fixed2csv
fi

if [ ! -e "qps2csv" ]; then
	go build -o qps2csv ./cmd/qps2csv
fi

./fixed2csv -i ./results/fixed/benchmarks.txt -d ./results/fixed
./qps2csv -i ./results/qps/benchmarks.txt -d ./results/qps

rm fixed2csv
rm qps2csv