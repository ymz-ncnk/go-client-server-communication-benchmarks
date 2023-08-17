package grpc

import (
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MakeClient(addr string) (EchoServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithWriteBufferSize(utils.IOBufSize),
		grpc.WithReadBufferSize(utils.IOBufSize),
	)

	if err != nil {
		return nil, nil, err
	}
	return NewEchoServiceClient(conn), conn, nil
}
