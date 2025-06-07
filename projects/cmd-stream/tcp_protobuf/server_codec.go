package tcpproto

import (
	"io"

	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/transport-go"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/projects/cmd-stream/tcp_protobuf/cmds"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/projects/cmd-stream/tcp_protobuf/receiver"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/projects/cmd-stream/tcp_protobuf/results"
	"google.golang.org/protobuf/proto"
)

func NewServerCodec(bs []byte) ServerCodec {
	return ServerCodec{bs}
}

type ServerCodec struct {
	bs []byte
}

func (c ServerCodec) Encode(result core.Result, w transport.Writer) (n int,
	err error) {
	bs, err := proto.Marshal(result.(results.EchoResult).ProtoData)
	if err != nil {
		return
	}
	n, err = varint.PositiveInt.Marshal(len(bs), w)
	if err != nil {
		return
	}
	n, err = w.Write(bs)
	return
}

func (c ServerCodec) Decode(r transport.Reader) (cmd core.Cmd[receiver.Receiver],
	n int, err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		n1 int
		bs = make([]byte, l)
	)
	n1, err = io.ReadFull(r, bs)
	n += n1
	if err != nil {
		return
	}
	data := common.ProtoData{}
	err = proto.Unmarshal(bs, &data)
	if err != nil {
		return
	}
	cmd = cmds.EchoCmd{ProtoData: &data}
	return
}
