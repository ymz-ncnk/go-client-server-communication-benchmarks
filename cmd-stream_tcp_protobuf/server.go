package cstcpb

import (
	"io"
	"sync"

	"github.com/cmd-stream/base-go"
	base_server "github.com/cmd-stream/base-go/server"
	cs_server "github.com/cmd-stream/cmd-stream-go/server"
	"github.com/cmd-stream/delegate-go"
	"github.com/cmd-stream/transport-go"
	transport_common "github.com/cmd-stream/transport-go/common"
	"github.com/mus-format/mus-stream-go/varint"
	data_protobuf "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
	"google.golang.org/protobuf/proto"
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
		ServerCodec{make([]byte, 2000)},
		struct{}{},
		nil)
}

type ServerCodec struct {
	bs []byte
}

func (c ServerCodec) Encode(result base.Result, w transport.Writer) (err error) {
	bs, err := proto.Marshal(result.(EchoCmd).Data)
	if err != nil {
		return
	}
	_, err = varint.MarshalPositiveInt(len(bs), w)
	if err != nil {
		return
	}
	_, err = w.Write(bs)
	return
}

func (c ServerCodec) Decode(r transport.Reader) (cmd base.Cmd[struct{}],
	err error) {
	l, _, err := varint.UnmarshalPositiveInt(r)
	if err != nil {
		return
	}
	bs := make([]byte, l)
	_, err = io.ReadFull(r, bs)
	if err != nil {
		return
	}
	data := data_protobuf.Data{}
	err = proto.Unmarshal(bs, &data)
	if err != nil {
		return
	}
	cmd = EchoCmd{&data}
	return
}

func (c ServerCodec) Size(result base.Result) (size int) {
	panic("unimplemented")
}
