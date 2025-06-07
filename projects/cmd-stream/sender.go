package cs

import (
	"errors"
	"net"
	"time"

	cmdstream "github.com/cmd-stream/cmd-stream-go"
	cln "github.com/cmd-stream/cmd-stream-go/client"
	grp "github.com/cmd-stream/cmd-stream-go/group"
	sndr "github.com/cmd-stream/sender-go"
	"github.com/cmd-stream/transport-go"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func MakeSender[T any](addr string, clientsCount int, codec cln.Codec[T]) (
	sender sndr.Sender[T], err error) {
	var connFactory cln.ConnFactoryFn = func() (net.Conn, error) {
		return net.Dial("tcp", addr)
	}
	group, err := cmdstream.MakeClientGroup(clientsCount, codec, connFactory,
		grp.WithClientOps[T](
			cln.WithTransport(
				transport.WithWriterBufSize(common.IOBufSize),
				transport.WithReaderBufSize(common.IOBufSize),
			),
		),
	)
	if err != nil {
		group.Close()
		return
	}
	sender = sndr.New(group)
	return
}

func CloseSender[T any](sender sndr.Sender[T]) (err error) {
	err = sender.Close()
	if err != nil {
		return
	}
	select {
	case <-time.NewTimer(time.Second).C:
		return errors.New("can't close the sender, cause timeout exceeded")
	case <-sender.Done():
		return
	}
}
