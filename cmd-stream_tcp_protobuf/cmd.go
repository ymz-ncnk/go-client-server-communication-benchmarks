package cstcpb

import (
	"context"
	"time"

	base "github.com/cmd-stream/base-go"
	data_protobuf "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

type EchoCmd struct {
	*data_protobuf.Data
}

func (cmd EchoCmd) Exec(ctx context.Context, at time.Time, seq base.Seq,
	receiver struct{},
	proxy base.Proxy,
) error {
	time.Sleep(utils.Delay)
	return proxy.Send(seq, cmd)
}

func (cmd EchoCmd) LastOne() bool {
	return true
}
