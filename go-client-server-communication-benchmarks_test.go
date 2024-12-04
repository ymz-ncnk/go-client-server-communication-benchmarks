package gcscb

import (
	"errors"
	"flag"
	"net"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cloudwego/kitex/server"
	"github.com/cmd-stream/base-go/client"
	base_server "github.com/cmd-stream/base-go/server"
	cstm "github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream_tcp_mus"
	cstp "github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream_tcp_protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/data"
	data_pb "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	grpc "github.com/ymz-ncnk/go-client-server-communication-benchmarks/grpc_http2_protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

const (
	NsOpMetric = "ns/op"
	NsMetric   = "ns"
	NMetric    = "N"
)

type ExchangeFn_gRPC = func(data *data_pb.Data,
	client grpc.EchoServiceClient, wg *sync.WaitGroup, b *testing.B)

type ExchangeFn_Kitex = func(data *echo.KitexData,
	client kitexechoservice.Client, wg *sync.WaitGroup, b *testing.B)

type ExchangeFn_CmdStream = func(cmd cstm.EchoCmd,
	client *client.Client[struct{}], wg *sync.WaitGroup, b *testing.B)

type ExchangeFn_CmdStream_Protobuf = func(cmd cstp.EchoCmd,
	client *client.Client[struct{}], wg *sync.WaitGroup, b *testing.B)

func BenchmarkQPS(b *testing.B) {
	var (
		ds      = data.GenerateDataSet(16, utils.GenSize())
		grpcDS  = ds.ToProtobuf()
		kitexDS = kitex.ConvertDataSet(ds)
		cstmDS  = cstm.ConvertDataSet(ds)
		cstpDS  = cstp.ConvertDataSet(ds)
	)

	b.Run("1", func(b *testing.B) {
		clientsCount := 1

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_MUS(clientsCount, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount, cstpDS, b)
		})
	})

	b.Run("2", func(b *testing.B) {
		clientsCount := 2

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_MUS(clientsCount, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount, cstpDS, b)
		})
	})

	b.Run("4", func(b *testing.B) {
		clientsCount := 4

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_MUS(clientsCount, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount, cstpDS, b)
		})
	})

	b.Run("8", func(b *testing.B) {
		clientsCount := 8

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_MUS(clientsCount, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount, cstpDS, b)
		})
	})

	b.Run("16", func(b *testing.B) {
		clientsCount := 16

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_MUS(clientsCount, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount, cstpDS, b)
		})
	})

}

func BenchmarkFixed(b *testing.B) {
	n, err := N()
	if err != nil {
		b.Fatal(err)
	}

	var (
		ds      = data.GenerateDataSet(16, n)
		grpcDS  = ds.ToProtobuf()
		kitexDS = kitex.ConvertDataSet(ds)
		cstmDS  = cstm.ConvertDataSet(ds)
		cstpDS  = cstp.ConvertDataSet(ds)
	)

	b.Run("1", func(b *testing.B) {
		clientsCount := 1

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n, cstpDS, b)
		})
	})

	b.Run("2", func(b *testing.B) {
		clientsCount := 2

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n, cstpDS, b)
		})
	})

	b.Run("4", func(b *testing.B) {
		clientsCount := 4

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n, cstpDS, b)
		})
	})

	b.Run("8", func(b *testing.B) {
		clientsCount := 8

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n, cstpDS, b)
		})
	})

	b.Run("16", func(b *testing.B) {
		clientsCount := 16

		b.Run("grpc_http2_protobuf", func(b *testing.B) {
			benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n, grpcDS, b)
		})
		b.Run("kitex_ttheader_protobuf", func(b *testing.B) {
			benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n, kitexDS, b)
		})
		b.Run("cmd-stream_tcp_mus", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n, cstmDS, b)
		})
		b.Run("cmd-stream_tcp_protobuf", func(b *testing.B) {
			benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n, cstpDS, b)
		})
	})

}

// -----------------------------------------------------------------------------
// gRPC/HTTP2/Protobuf
// -----------------------------------------------------------------------------

func benchmarkQPS_gRPC_HTTP2_Protobuf(clientsCount int,
	ds data_pb.DataSetProtobuf, b *testing.B) {
	benchmark_gRPC_HTTP2_Protobuf(clientsCount, 0, ds, grpc.ExchangeQPS, b)
	b.ReportMetric(0, NsOpMetric)
	b.ReportMetric(float64(b.Elapsed()), NsMetric)
}

func benchmarkFixed_gRPC_HTTP2_Protobuf(clientsCount, n int,
	ds data_pb.DataSetProtobuf, b *testing.B) {
	var (
		copsD      = make(chan time.Duration, n)
		exchangeFn = func(data *data_pb.Data, client grpc.EchoServiceClient,
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			grpc.ExchangeFixed(data, client, copsD, wg, b)
		}
		N = n / clientsCount
	)
	benchmark_gRPC_HTTP2_Protobuf(clientsCount, N, ds, exchangeFn, b)
	b.ReportMetric(float64(N), NMetric)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmark_gRPC_HTTP2_Protobuf(clientsCount, N int,
	ds data_pb.DataSetProtobuf,
	exchangeFn ExchangeFn_gRPC,
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9001"
		l    net.Listener
		err  error
		wgS  = &sync.WaitGroup{}
	)
	if l, err = startServer_gRPC(addr, wgS); err != nil {
		b.Fatal(err)
	}
	clients, err := makeClients_gRPC(addr, clientsCount)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	wg := &sync.WaitGroup{}
	var i int
	for i = 0; i < b.N; i++ {
		if N != 0 && i == N {
			break
		}
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			go exchangeFn(ds[j][i], clients[j], wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	if err = stopServer_gRPC(l, wgS); err != nil {
		b.Fatal(err)
	}
}

func startServer_gRPC(addr string, wg *sync.WaitGroup) (l net.Listener, err error) {
	if l, err = net.Listen("tcp", addr); err != nil {
		return
	}
	wg.Add(1)
	grpc.StartServer(l, wg)
	return
}

func makeClients_gRPC(addr string, count int) (clients []grpc.EchoServiceClient,
	err error) {
	clients = make([]grpc.EchoServiceClient, count)
	for i := 0; i < count; i++ {
		clients[i], _, err = grpc.MakeClient(addr)
		if err != nil {
			return
		}
	}
	return
}

func stopServer_gRPC(l net.Listener, wg *sync.WaitGroup) (err error) {
	if err = l.Close(); err != nil {
		return
	}
	wg.Wait()
	return
}

// -----------------------------------------------------------------------------
// Kitex/TTHeader/Protobuf
// -----------------------------------------------------------------------------

func benchmarkQPS_Kitex_TTHeader_Protobuf(clientsCount int,
	ds kitex.DataSet,
	b *testing.B,
) {
	benchmark_Kitex_TTHeader_Protobuf(clientsCount, 0, ds, kitex.ExchangeQPS, b)
	b.ReportMetric(0, NsOpMetric)
	b.ReportMetric(float64(b.Elapsed()), NsMetric)
}

func benchmarkFixed_Kitex_TTHeader_Protobuf(clientsCount, n int,
	ds kitex.DataSet,
	b *testing.B,
) {
	var (
		copsD                       = make(chan time.Duration, n)
		exchangeFn ExchangeFn_Kitex = func(data *echo.KitexData,
			client kitexechoservice.Client,
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			kitex.ExchangeFixed(data, client, copsD, wg, b)
		}
		N = n / clientsCount
	)
	benchmark_Kitex_TTHeader_Protobuf(clientsCount, N, ds,
		exchangeFn, b)
	b.ReportMetric(float64(N), NMetric)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmark_Kitex_TTHeader_Protobuf(clientsCount, N int,
	ds kitex.DataSet,
	exchangeFn ExchangeFn_Kitex,
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9002"
		wgS  = &sync.WaitGroup{}
	)
	srv := startServer_Kitex(addr, wgS)
	clients, err := makeClients_Kitex(addr, clientsCount)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		if N != 0 && i == N {
			break
		}
		wg.Add(len(clients))
		for j := 0; j < len(clients); j++ {
			go exchangeFn(ds[j][i], clients[j], wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	if err = stopServer_Kitex(srv, wgS); err != nil {
		b.Fatal(err)
	}
}

func startServer_Kitex(addr string, wg *sync.WaitGroup) (srv server.Server) {
	wg.Add(1)
	return kitex.StartServer(addr, wg)
}

func makeClients_Kitex(addr string, count int) (clients []kitexechoservice.Client,
	err error) {
	clients = make([]kitexechoservice.Client, count)
	for i := 0; i < count; i++ {
		clients[i], err = kitex.MakeClient(addr)
		if err != nil {
			return
		}
	}
	return
}

func stopServer_Kitex(srv server.Server, wg *sync.WaitGroup) (err error) {
	if err = srv.Stop(); err != nil {
		return
	}
	wg.Wait()
	return
}

// -----------------------------------------------------------------------------
// cmd-stream/TCP/MUS
// -----------------------------------------------------------------------------

func benchmarkQPS_CmdStream_TCP_MUS(clientsCount int,
	ds cstm.DataSet,
	b *testing.B,
) {
	benchmark_CmdStream_TCP_MUS(clientsCount, 0, ds, cstm.ExchangeQPS, b)
	b.ReportMetric(0, NsOpMetric)
	b.ReportMetric(float64(b.Elapsed()), NsMetric)
}

func benchmarkFixed_CmdStream_TCP_MUS(clientsCount, n int,
	ds cstm.DataSet,
	b *testing.B,
) {
	var (
		copsD      = make(chan time.Duration, n)
		exchangeFn = func(cmd cstm.EchoCmd, client *client.Client[struct{}],
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			cstm.ExchangeFixed(cmd, client, copsD, wg, b)
		}
		N = n / clientsCount
	)
	benchmark_CmdStream_TCP_MUS(clientsCount, N, ds, exchangeFn, b)
	b.ReportMetric(float64(N), NMetric)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmark_CmdStream_TCP_MUS(clientsCount, N int,
	ds cstm.DataSet,
	exchangFn ExchangeFn_CmdStream,
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9003"
		wgS  = &sync.WaitGroup{}
	)
	srv, err := startServer_CmdStream(addr, clientsCount, wgS)
	if err != nil {
		b.Fatal(err)
	}
	clients, err := makeClients_CmdStream(addr, clientsCount)
	if err != nil {
		return
	}

	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		if N != 0 && i == N {
			break
		}
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			go exchangFn(ds[j][i], clients[j], wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	if err = stopServer_CmdStream(srv, wgS); err != nil {
		b.Fatal(err)
	}
}

func startServer_CmdStream(addr string, clientsCount int, wg *sync.WaitGroup) (
	srv *base_server.Server, err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	wg.Add(1)
	if srv, err = cstm.StartServer(clientsCount, l.(*net.TCPListener), wg); err != nil {
		return
	}
	return
}

func makeClients_CmdStream(addr string, count int) (
	clients []*client.Client[struct{}], err error) {
	var conn net.Conn
	clients = make([]*client.Client[struct{}], count)
	for i := 0; i < count; i++ {
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			return
		}
		clients[i], err = cstm.MakeClient(conn)
		if err != nil {
			return
		}
	}
	return
}

func stopServer_CmdStream(srv *base_server.Server, wg *sync.WaitGroup) (
	err error) {
	if err = srv.Close(); err != nil {
		return
	}
	wg.Wait()
	return
}

// -----------------------------------------------------------------------------
// cmd-stream/TCP/Protobuf
// -----------------------------------------------------------------------------

// If you are looking for an example of using cmd-stream/Protobuf also check
// https://github.com/cmd-stream/cmd-stream-examples-go/tree/main/standard_protobuf.

func benchmarkQPS_CmdStream_TCP_Protobuf(clientsCount int,
	ds cstp.DataSet,
	b *testing.B,
) {
	benchmark_CmdStream_TCP_Protobuf(clientsCount, 0, ds, cstp.ExchangeQPS, b)
	b.ReportMetric(0, NsOpMetric)
	b.ReportMetric(float64(b.Elapsed()), NsMetric)
}

func benchmarkFixed_CmdStream_TCP_Protobuf(clientsCount, n int,
	ds cstp.DataSet,
	b *testing.B,
) {
	var (
		copsD      = make(chan time.Duration, n)
		exchangeFn = func(cmd cstp.EchoCmd, client *client.Client[struct{}],
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			cstp.ExchangeFixed(cmd, client, copsD, wg, b)
		}
		N = n / clientsCount
	)
	benchmark_CmdStream_TCP_Protobuf(clientsCount, N, ds, exchangeFn, b)
	b.ReportMetric(float64(N), NMetric)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmark_CmdStream_TCP_Protobuf(clientsCount, N int,
	ds cstp.DataSet,
	exchangFn ExchangeFn_CmdStream_Protobuf,
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9004"
		wgS  = &sync.WaitGroup{}
	)
	srv, err := startServer_CmdStream_Protobuf(addr, clientsCount, wgS)
	if err != nil {
		b.Fatal(err)
	}
	clients, err := makeClients_CmdStream_Protobuf(addr, clientsCount)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		if N != 0 && i == N {
			break
		}
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			go exchangFn(ds[j][i], clients[j], wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	if err = stopServer_CmdStream_Protobuf(srv, wgS); err != nil {
		b.Fatal(err)
	}
}

func startServer_CmdStream_Protobuf(addr string, clientsCount int, wg *sync.WaitGroup) (
	srv *base_server.Server, err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	wg.Add(1)
	if srv, err = cstp.StartServer(clientsCount, l.(*net.TCPListener), wg); err != nil {
		return
	}
	return
}

func makeClients_CmdStream_Protobuf(addr string, count int) (
	clients []*client.Client[struct{}], err error) {
	var conn net.Conn
	clients = make([]*client.Client[struct{}], count)
	for i := 0; i < count; i++ {
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			return
		}
		clients[i], err = cstp.MakeClient(conn)
		if err != nil {
			return
		}
	}
	return
}

func stopServer_CmdStream_Protobuf(srv *base_server.Server, wg *sync.WaitGroup) (
	err error) {
	if err = srv.Close(); err != nil {
		return
	}
	wg.Wait()
	return
}

func N() (N int, err error) {
	var (
		f         = flag.Lookup("test.benchtime")
		benchtime = f.Value.String()
	)
	if !strings.HasSuffix(benchtime, "x") {
		err = errors.New("you should specify -benchtime with x suffix")
		return
	}
	return strconv.Atoi(benchtime[:len(benchtime)-1])
}
