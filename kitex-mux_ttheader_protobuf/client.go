package kitex

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
)

func MakeClient(addr string) (kitexechoservice.Client, error) {
	var opts []client.Option
	opts = append(opts, client.WithHostPorts(addr))
	opts = append(opts, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	opts = append(opts, client.WithMuxConnection(1))
	return kitexechoservice.NewClient("echo", opts...)
}
