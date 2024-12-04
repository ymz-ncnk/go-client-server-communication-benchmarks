package grpc

import (
	context "context"
	"sync"
	"testing"
	"time"

	data_protobuf "github.com/ymz-ncnk/go-client-server-communication-benchmarks/data/protobuf"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

func ExchangeFixed(data *data_protobuf.Data, client EchoServiceClient,
	copsD chan<- time.Duration,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	start := time.Now()
	r, err := exchange(data, client)
	if err != nil {
		b.Error(err)
		return
	}
	utils.QueueCopD(copsD, time.Since(start))
	if !data_protobuf.EqualData(data, r) {
		b.Error("unexpected result")
	}
}

func ExchangeQPS(data *data_protobuf.Data, client EchoServiceClient,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	r, err := exchange(data, client)
	if err != nil {
		b.Error(err)
		return
	}
	if !data_protobuf.EqualData(data, r) {
		b.Error("unexpected result")
	}
}

func exchange(data *data_protobuf.Data, client EchoServiceClient,
) (r *data_protobuf.Data, err error) {
	return client.Echo(context.Background(), data)
}
