package csp

import (
	"io"
	"net"

	base "github.com/cmd-stream/base-go"
	base_client "github.com/cmd-stream/base-go/client"
	cs_client "github.com/cmd-stream/cmd-stream-go/client"
	cs_server "github.com/cmd-stream/cmd-stream-go/server"
	"github.com/cmd-stream/transport-go"
	transport_common "github.com/cmd-stream/transport-go/common"
	"github.com/mus-format/mus-stream-go/varint"
	data_protobuf "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
	"google.golang.org/protobuf/proto"
)

func MakeClient(conn net.Conn) (client *base_client.Client[struct{}],
	err error) {
	conf := cs_client.Conf{
		Transport: transport_common.Conf{
			WriterBufSize: utils.IOBufSize,
			ReaderBufSize: utils.IOBufSize,
		},
	}
	return cs_client.New[struct{}](cs_server.DefServerInfo, conf,
		ClientCodec{bs: make([]byte, 2000)}, conn, nil)
}

type ClientCodec struct {
	bs []byte
}

func (c ClientCodec) Encode(cmd base.Cmd[struct{}], w transport.Writer) (
	err error) {
	bs, err := proto.Marshal(cmd.(EchoCmd).Data)
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

func (c ClientCodec) Decode(r transport.Reader) (result base.Result,
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
	result = EchoCmd{&data}
	return
}

func (codec ClientCodec) Size(cmd base.Cmd[struct{}]) (size int) {
	panic("unimplemented")
}
