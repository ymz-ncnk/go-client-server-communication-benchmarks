package cstcpb

import (
	"net"

	base "github.com/cmd-stream/base-go"
	base_client "github.com/cmd-stream/base-go/client"
	cs_client "github.com/cmd-stream/cmd-stream-go/client"
	"github.com/cmd-stream/transport-go"
	data_mus "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/data/muss"
)

func MakeClient(conn net.Conn) (client *base_client.Client[struct{}],
	err error) {
	return cs_client.NewDef[struct{}](ClientCodec{}, conn, nil)
}

type ClientCodec struct{}

func (c ClientCodec) Encode(cmd base.Cmd[struct{}],
	w transport.Writer) (err error) {
	_, err = data_mus.MarshalDataMUS(data_mus.Data(cmd.(EchoCmd)), w)
	return
}

func (c ClientCodec) Decode(r transport.Reader) (result base.Result,
	err error) {
	d, _, err := data_mus.UnmarshalDataMUS(r)
	result = EchoCmd(d)
	return
}

func (codec ClientCodec) Size(cmd base.Cmd[struct{}]) (size int) {
	return data_mus.SizeDataMUS(data_mus.Data(cmd.(EchoCmd)))
}
