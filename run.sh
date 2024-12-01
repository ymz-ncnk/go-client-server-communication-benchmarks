#!/bin/bash

readonly QPS_GEN_SIZE=400000
readonly FIXED_GEN_SIZE=100000

clients=(1 2 4 8 16)
run_qps() {
	for i in "${clients[@]}"
	do
		echo "Run $1 for $i client/s."
		s=$(expr $QPS_GEN_SIZE / $i)
		CLIENTS_COUNT="$i" GEN_SIZE="$s" go test -bench="$1" \
			-count=10 \
			> ./results/qps/"$2"_"$i".txt
	done	
}

run_fixed () {
	for i in "${clients[@]}"
	do
		echo "Run $1 for $i client/s."
		s=$(expr $FIXED_GEN_SIZE / $i) 
		CLIENTS_COUNT="$i" GEN_SIZE="$s" go test -bench="$1" \
			-benchtime="$s"x \
			-benchmem \
			-count=10 \
			> ./results/fixed/"$2"_"$i".txt
	done
}

run_qps "BenchmarkQPS_GRPC_HTTP2_Protobuf" "grpc_http2_protobuf"
run_qps "BenchmarkQPS_Kitex_TTHeader_Protobuf" "kitex-mux_ttheader_protobuf"
run_qps "BenchmarkQPS_CmdStream_TCP_Protobuf" "cmd-stream_tcp_protobuf"
run_qps "BenchmarkQPS_CmdStream_TCP_MUS" "cmd-stream_tcp_mus"

run_fixed "BenchmarkFixed_GRPC_HTTP2_Protobuf" "grpc_http2_protobuf"
run_fixed "BenchmarkFixed_Kitex_TTHeader_Protobuf" "kitex-mux_ttheader_protobuf"
run_fixed "BenchmarkFixed_CmdStream_TCP_Protobuf" "cmd-stream_tcp_protobuf"
run_fixed "BenchmarkFixed_CmdStream_TCP_MUS" "cmd-stream_tcp_mus"
