package kitex

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	server "github.com/cloudwego/kitex/server"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
)

func init() {
	klog.SetLevel(klog.LevelError)
}

type EchoImpl struct{}

func (s *EchoImpl) Echo(ctx context.Context, req *echo.KitexData) (resp *echo.KitexData, err error) {
	time.Sleep(utils.Delay)
	return &echo.KitexData{
		Bool:    req.Bool,
		Int64:   req.Int64,
		String_: req.String_,
		Float64: req.Float64,
	}, nil
}

func StartServer(addr string, wg *sync.WaitGroup) (srv server.Server) {
	a, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(a))
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	opts = append(opts, server.WithMuxTransport())
	srv = kitexechoservice.NewServer(new(EchoImpl), opts...)
	go func() {
		defer wg.Done()
		srv.Run()
	}()
	return
}
