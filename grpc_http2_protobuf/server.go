package grpc

import (
	context "context"
	"net"
	"sync"
	"time"

	data_protobuf "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
	grpc "google.golang.org/grpc"
)

type echoServer struct {
	UnimplementedEchoServiceServer
}

func (server echoServer) Echo(ctx context.Context, data *data_protobuf.Data) (
	*data_protobuf.Data, error) {
	time.Sleep(utils.Delay)
	return &data_protobuf.Data{
		Bool:    data.Bool,
		Int64:   data.Int64,
		String_: data.String_,
		Float64: data.Float64,
	}, nil
}

func StartServer(l net.Listener, wg *sync.WaitGroup) {
	server := MakeServer()
	go func() {
		defer wg.Done()
		server.Serve(l)
	}()
}

func MakeServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.WriteBufferSize(utils.IOBufSize),
		grpc.ReadBufferSize(utils.IOBufSize),
	}
	grpcServer := grpc.NewServer(opts...)
	RegisterEchoServiceServer(grpcServer, echoServer{})
	return grpcServer
}
