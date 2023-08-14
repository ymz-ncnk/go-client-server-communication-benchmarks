package cstcpb

import (
	"context"
	"time"

	base "github.com/cmd-stream/base-go"
	data_mus "github.com/ymz-ncnk/go-inter-server-communication-benchmarks/data/muss"
	"github.com/ymz-ncnk/go-inter-server-communication-benchmarks/utils"
)

type EchoCmd data_mus.Data

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
