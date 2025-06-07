package http2proto

import (
	context "context"
	"net"
	"sync"
	"time"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
	grpc "google.golang.org/grpc"
)

func StartServer(addr string, wg *sync.WaitGroup) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return
	}
	wg.Add(1)
	server := makeServer()
	go func() {
		defer wg.Done()
		server.Serve(listener)
	}()
	return
}

func CloseServer(l net.Listener, wg *sync.WaitGroup) (err error) {
	if err = l.Close(); err != nil {
		return
	}
	wg.Wait()
	return
}

func makeServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.WriteBufferSize(common.IOBufSize),
		grpc.ReadBufferSize(common.IOBufSize),
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterEchoServiceServer(grpcServer, echoServer{})
	return grpcServer
}

type echoServer struct {
	UnimplementedEchoServiceServer
}

func (server echoServer) Echo(ctx context.Context, data *common.ProtoData) (
	*common.ProtoData, error) {
	time.Sleep(common.Delay)
	return &common.ProtoData{
		Bool:    data.Bool,
		Int64:   data.Int64,
		String_: data.String_,
		Float64: data.Float64,
	}, nil
}
