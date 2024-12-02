package gcscb

import (
	"math/rand"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cmd-stream/base-go/client"
	cs "github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream_tcp_mus"
	csp "github.com/ymz-ncnk/go-client-server-communication-benchmarks/cmd-stream_tcp_protobuf"
	data_mus "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/mus"
	data_protobuf "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	grpc "github.com/ymz-ncnk/go-client-server-communication-benchmarks/grpc_http2_protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

const CopsDChanSize = 5000000

// -----------------------------------------------------------------------------
func BenchmarkQPS_GRPC_HTTP2_Protobuf(b *testing.B) {
	clientsCount := utils.ClientsCount()
	benchmarkGRPC_HTTP2_Protobuf(clientsCount, grpc.ExchangeQPS, b)
	b.ReportMetric(0, "ns/op")
	b.ReportMetric(float64(b.Elapsed()), "ns")
}

func BenchmarkFixed_GRPC_HTTP2_Protobuf(b *testing.B) {
	var (
		clientsCount = utils.ClientsCount()
		copsD        = make(chan time.Duration, CopsDChanSize)
		exchangeFn   = func(data *data_protobuf.Data, client grpc.EchoServiceClient,
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			grpc.ExchangeFixed(data, client, copsD, wg, b)
		}
	)
	benchmarkGRPC_HTTP2_Protobuf(clientsCount, exchangeFn, b)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmarkGRPC_HTTP2_Protobuf(clientsCount int,
	exchangeFn func(data *data_protobuf.Data, client grpc.EchoServiceClient, wg *sync.WaitGroup, b *testing.B),
	b *testing.B) {
	var (
		addr = "127.0.0.1:9001"
		r    = rand.New(rand.NewSource(time.Now().Unix()))
		ds   = data_protobuf.GenData(clientsCount, utils.GenSize(), r)
		wgS  = &sync.WaitGroup{}
	)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	wgS.Add(1)
	grpc.StartServer(l, wgS)
	clients := make([]grpc.EchoServiceClient, clientsCount)
	for i := 0; i < clientsCount; i++ {
		clients[i], _, err = grpc.MakeClient(addr)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			data := ds[j][i]
			client := clients[j]
			go exchangeFn(data, client, wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	err = l.Close()
	if err != nil {
		b.Fatal(err)
	}
	wgS.Wait()
}

// -----------------------------------------------------------------------------
func BenchmarkQPS_Kitex_TTHeader_Protobuf(b *testing.B) {
	clientsCount := utils.ClientsCount()
	benchmarkKitex_TTHeader_Protobuf(clientsCount, kitex.ExchangeQPS, b)
	b.ReportMetric(0, "ns/op")
	b.ReportMetric(float64(b.Elapsed()), "ns")
}

func BenchmarkFixed_Kitex_TTHeader_Protobuf(b *testing.B) {
	var (
		clientsCount = utils.ClientsCount()
		copsD        = make(chan time.Duration, CopsDChanSize)
		exchangeFn   = func(data *echo.KitexData, client kitexechoservice.Client,
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			kitex.ExchangeFixed(data, client, copsD, wg, b)
		}
	)
	benchmarkKitex_TTHeader_Protobuf(clientsCount, exchangeFn, b)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmarkKitex_TTHeader_Protobuf(clientsCount int,
	exchangeFn func(data *echo.KitexData, client kitexechoservice.Client, wg *sync.WaitGroup, b *testing.B),
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9002"
		r    = rand.New(rand.NewSource(time.Now().Unix()))
		ds   = kitex.GenData(clientsCount, utils.GenSize(), r)
		err  error
		wgS  = &sync.WaitGroup{}
	)
	wgS.Add(1)
	srv := kitex.StartServer(addr, wgS)
	clients := make([]kitexechoservice.Client, clientsCount)
	for i := 0; i < clientsCount; i++ {
		clients[i], err = kitex.MakeClient(addr)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(len(clients))
		for j := 0; j < len(clients); j++ {
			data := ds[j][i]
			client := clients[j]
			go exchangeFn(data, client, wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()

	err = srv.Stop()
	if err != nil {
		b.Fatal(err)
	}
	wgS.Wait()
}

// -----------------------------------------------------------------------------
func BenchmarkQPS_CmdStream_TCP_MUS(b *testing.B) {
	clientsCount := utils.ClientsCount()
	benchmarkCmdStream_TCP_MUS(clientsCount, cs.ExchangeQPS, b)
	b.ReportMetric(0, "ns/op")
	b.ReportMetric(float64(b.Elapsed()), "ns")
}

func BenchmarkFixed_CmdStream_TCP_MUS(b *testing.B) {
	var (
		clientsCount = utils.ClientsCount()
		copsD        = make(chan time.Duration, CopsDChanSize)
		exchangeFn   = func(cmd cs.EchoCmd, client *client.Client[struct{}],
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			cs.ExchangeFixed(cmd, client, copsD, wg, b)
		}
	)
	benchmarkCmdStream_TCP_MUS(clientsCount, exchangeFn, b)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmarkCmdStream_TCP_MUS(clientsCount int,
	exchangFn func(cmd cs.EchoCmd, client *client.Client[struct{}], wg *sync.WaitGroup, b *testing.B),
	b *testing.B,
) {
	var (
		addr = "127.0.0.1:9003"
		r    = rand.New(rand.NewSource(time.Now().Unix()))
		ds   = data_mus.GenData(clientsCount, utils.GenSize(), r)
		wgS  = &sync.WaitGroup{}
	)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		b.Fatal(err)
	}
	wgS.Add(1)
	server, err := cs.StartServer(clientsCount, listener.(*net.TCPListener), wgS)
	if err != nil {
		b.Fatal(err)
	}
	clients := make([]*client.Client[struct{}], clientsCount)
	for i := 0; i < clientsCount; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			b.Fatal(err)
		}
		clients[i], err = cs.MakeClient(conn)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			data := ds[j][i]
			client := clients[j]
			go exchangFn(cs.EchoCmd(data), client, wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()
	err = server.Close()
	if err != nil {
		b.Fatal(err)
	}
	wgS.Wait()
}

// -----------------------------------------------------------------------------
func BenchmarkQPS_CmdStream_TCP_Protobuf(b *testing.B) {
	clientsCount := utils.ClientsCount()
	benchmarkCmdStream_TCP_Protobuf(clientsCount, csp.ExchangeQPS, b)
	b.ReportMetric(0, "ns/op")
	b.ReportMetric(float64(b.Elapsed()), "ns")
}

func BenchmarkFixed_CmdStream_TCP_Protobuf(b *testing.B) {
	var (
		clientsCount = utils.ClientsCount()
		copsD        = make(chan time.Duration, CopsDChanSize)
		exchangeFn   = func(cmd csp.EchoCmd, client *client.Client[struct{}],
			wg *sync.WaitGroup,
			b *testing.B,
		) {
			csp.ExchangeFixed(cmd, client, copsD, wg, b)
		}
	)
	benchmarkCmdStream_TCP_Protobuf(clientsCount, exchangeFn, b)
	utils.ReportMetrics(clientsCount, copsD, b)
}

func benchmarkCmdStream_TCP_Protobuf(clientsCount int,
	exchangFn func(cmd csp.EchoCmd, client *client.Client[struct{}], wg *sync.WaitGroup, b *testing.B),
	b *testing.B,
) {
	// If you are looking for an example of using cmd-stream/Protobuf also have a
	// look at https://github.com/cmd-stream/cmd-stream-examples-go/tree/main/standard_protobuf.
	var (
		addr = "127.0.0.1:9004"
		r    = rand.New(rand.NewSource(time.Now().Unix()))
		ds   = data_protobuf.GenData(clientsCount, utils.GenSize(), r)
		wgS  = &sync.WaitGroup{}
	)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		b.Fatal(err)
	}
	wgS.Add(1)
	server, err := csp.StartServer(clientsCount, listener.(*net.TCPListener), wgS)
	if err != nil {
		b.Fatal(err)
	}
	clients := make([]*client.Client[struct{}], clientsCount)
	for i := 0; i < clientsCount; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			b.Fatal(err)
		}
		clients[i], err = csp.MakeClient(conn)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(clientsCount)
		for j := 0; j < clientsCount; j++ {
			data := ds[j][i]
			client := clients[j]
			go exchangFn(csp.EchoCmd{Data: data}, client, wg, b)
		}
	}
	wg.Wait()
	b.StopTimer()
	err = server.Close()
	if err != nil {
		b.Fatal(err)
	}
	wgS.Wait()
}
