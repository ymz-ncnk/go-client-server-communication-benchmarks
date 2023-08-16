package cstcpb

import (
	"sync"

	"github.com/cmd-stream/base-go"
	base_server "github.com/cmd-stream/base-go/server"
	cs_server "github.com/cmd-stream/cmd-stream-go/server"
	"github.com/cmd-stream/delegate-go"
	"github.com/cmd-stream/transport-go"
	transport_common "github.com/cmd-stream/transport-go/common"
	data_mus "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/data/mus"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
)

func StartServer(clientsCount int, l base.Listener, wg *sync.WaitGroup) (
	server *base_server.Server, err error) {
	server = MakeServer(clientsCount)
	go func() {
		defer wg.Done()
		server.Serve(l)
	}()
	return
}

func MakeServer(clientsCount int) *base_server.Server {
	conf := cs_server.Conf{
		Base: base_server.Conf{
			WorkersCount: clientsCount,
		},
		Transport: transport_common.Conf{
			WriterBufSize: utils.IOBufSize,
			ReaderBufSize: utils.IOBufSize,
		},
	}
	return cs_server.New[struct{}](cs_server.DefServerInfo,
		delegate.ServerSettings{},
		conf,
		ServerCodec{},
		struct{}{},
		nil)
}

type ServerCodec struct{}

func (c ServerCodec) Encode(result base.Result, w transport.Writer) (err error) {
	_, err = data_mus.MarshalDataMUS(data_mus.Data(result.(EchoCmd)), w)
	return
}

func (c ServerCodec) Decode(r transport.Reader) (cmd base.Cmd[struct{}],
	err error) {
	d, _, err := data_mus.UnmarshalDataMUS(r)
	cmd = EchoCmd(d)
	return
}

func (c ServerCodec) Size(result base.Result) (size int) {
	return data_mus.SizeDataMUS(data_mus.Data(result.(EchoCmd)))
}
