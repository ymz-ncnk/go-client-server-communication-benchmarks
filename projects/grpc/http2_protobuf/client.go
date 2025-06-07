package http2proto

import (
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MakeClients(addr string, count int) (clients []EchoServiceClient,
	err error) {
	clients = make([]EchoServiceClient, count)
	for i := range count {
		clients[i], _, err = makeClient(addr)
		if err != nil {
			return
		}
	}
	return
}

func makeClient(addr string) (EchoServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithWriteBufferSize(common.IOBufSize),
		grpc.WithReadBufferSize(common.IOBufSize),
	)
	if err != nil {
		return nil, nil, err
	}
	return NewEchoServiceClient(conn), conn, nil
}
