package kitex

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	kitex_utils "github.com/cloudwego/kitex/pkg/utils"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
)

func WithClientIOBufferSize() client.Option {
	return client.Option{F: func(o *client.Options, di *kitex_utils.Slice) {
		err := o.Configs.(rpcinfo.MutableRPCConfig).SetIOBufferSize(utils.IOBufSize)
		if err != nil {
			panic(err)
		}
	}}
}

func MakeClient(addr string) (kitexechoservice.Client, error) {
	var opts []client.Option
	opts = append(opts, client.WithHostPorts(addr))
	opts = append(opts, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	opts = append(opts, client.WithMuxConnection(1))
	opts = append(opts, WithClientIOBufferSize())
	return kitexechoservice.NewClient("echo", opts...)
}
