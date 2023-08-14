package cstcpb

import (
	"sync"

	"github.com/cmd-stream/base-go"
	base_server "github.com/cmd-stream/base-go/server"
	cs_server "github.com/cmd-stream/cmd-stream-go/server"
	"github.com/cmd-stream/delegate-go"
	"github.com/cmd-stream/transport-go"
	data_mus "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/data/muss"
)

func StartServer(clientsCount int, l base.Listener, wg *sync.WaitGroup) (
	server *base_server.Server, err error) {
	conf := cs_server.Conf{
		Base: base_server.Conf{WorkersCount: clientsCount},
	}
	server = cs_server.New[struct{}](cs_server.DefServerInfo,
		delegate.ServerSettings{},
		conf,
		ServerCodec{},
		struct{}{},
		nil)
	go func() {
		defer wg.Done()
		server.Serve(l)
	}()
	return
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
