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

func NewClientCodec(bs []byte) ClientCodec {
	return ClientCodec{bs}
}

type ClientCodec struct {
	bs []byte
}

func (c ClientCodec) Encode(cmd core.Cmd[receiver.Receiver], w transport.Writer) (
	n int, err error) {
	bs, err := proto.Marshal(cmd.(cmds.EchoCmd).ProtoData)
	if err != nil {
		return
	}
	n, err = varint.PositiveInt.Marshal(len(bs), w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = w.Write(bs)
	n += n1
	return
}

func (c ClientCodec) Decode(r transport.Reader) (result core.Result, n int,
	err error) {
	l, n, err := varint.PositiveInt.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		bs = make([]byte, l)
		n1 int
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
	result = results.EchoResult{ProtoData: &data}
	return
}
