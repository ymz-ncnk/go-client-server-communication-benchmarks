package cs

import (
	"net"
	"sync"

	cmdstream "github.com/cmd-stream/cmd-stream-go"
	srv "github.com/cmd-stream/cmd-stream-go/server"
	csrv "github.com/cmd-stream/core-go/server"
	"github.com/cmd-stream/transport-go"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func StartServer[T any](addr string, workersCount int, codec srv.Codec[T],
	receiver T, wg *sync.WaitGroup) (server *csrv.Server, err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	server = cmdstream.MakeServer(codec, srv.NewInvoker(receiver),
		srv.WithCore(
			csrv.WithWorkersCount(workersCount),
		),
		srv.WithTransport(
			transport.WithWriterBufSize(common.IOBufSize),
			transport.WithReaderBufSize(common.IOBufSize),
		),
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Serve(listener.(*net.TCPListener))
	}()
	return
}

func CloseServer(server *csrv.Server, wg *sync.WaitGroup) (err error) {
	err = server.Close()
	if err != nil {
		return
	}
	wg.Wait()
	return
}
