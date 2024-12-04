package cstm

import (
	"net"

	base "github.com/cmd-stream/base-go"
	base_client "github.com/cmd-stream/base-go/client"
	cs_client "github.com/cmd-stream/cmd-stream-go/client"
	cs_server "github.com/cmd-stream/cmd-stream-go/server"
	"github.com/cmd-stream/transport-go"
	transport_common "github.com/cmd-stream/transport-go/common"
	data_mus "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/mus"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

func MakeClient(conn net.Conn) (client *base_client.Client[struct{}],
	err error) {
	conf := cs_client.Conf{
		Transport: transport_common.Conf{
			WriterBufSize: utils.IOBufSize,
			ReaderBufSize: utils.IOBufSize,
		},
	}
	return cs_client.New[struct{}](cs_server.DefServerInfo, conf, ClientCodec{},
		conn, nil)
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
