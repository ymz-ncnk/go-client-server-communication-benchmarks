package cstcpb

import (
	"sync"
	"testing"
	"time"

	base "github.com/cmd-stream/base-go"
	base_client "github.com/cmd-stream/base-go/client"

	data_mus "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/mus"
)

func ExchangeQPS(cmd EchoCmd, client *base_client.Client[struct{}],
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	r, err := exchange(cmd, client)
	if err != nil {
		b.Error(err)
		return
	}
	if !data_mus.EqualData(data_mus.Data(cmd), data_mus.Data(r.Result.(EchoCmd))) {
		b.Error("unexpected result")
	}
}

func ExchangeFixed(cmd EchoCmd, client *base_client.Client[struct{}],
	copsD chan<- time.Duration,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	start := time.Now()
	r, err := exchange(cmd, client)
	if err != nil {
		b.Error(err)
		return
	}
	queueCopD(copsD, time.Since(start))
	if !data_mus.EqualData(data_mus.Data(cmd), data_mus.Data(r.Result.(EchoCmd))) {
		b.Error("unexpected result")
	}
}

func exchange(cmd EchoCmd,
	client *base_client.Client[struct{}]) (r base.AsyncResult, err error) {
	results := make(chan base.AsyncResult, 1)
	_, err = client.Send(cmd, results)
	if err != nil {
		return
	}
	r = <-results
	if r.Error != nil {
		err = r.Error
		return
	}
	return
}

func queueCopD(copsD chan<- time.Duration, spent time.Duration) {
	select {
	case copsD <- spent:
	default:
		panic("you should make the copsD channel bigger")
	}
}
