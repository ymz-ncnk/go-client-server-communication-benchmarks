package kitex

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/kitex-mux_ttheader_protobuf/kitex_gen/echo/kitexechoservice"
	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/utils"
)

func ExchangeFixed(data *echo.KitexData, client kitexechoservice.Client,
	copsD chan<- time.Duration,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	start := time.Now()
	r, err := exchange(data, client)
	if err != nil {
		b.Error(err)
	}
	utils.QueueCopD(copsD, time.Since(start))
	if !EqualData(data, r) {
		b.Error("unexpected result")
	}
}

func ExchangeQPS(data *echo.KitexData, client kitexechoservice.Client,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	r, err := exchange(data, client)
	if err != nil {
		b.Error(err)
	}
	if !EqualData(data, r) {
		b.Error("unexpected result")
	}
}

func exchange(data *echo.KitexData, client kitexechoservice.Client,
) (r *echo.KitexData, err error) {
	return client.Echo(context.Background(), data)
}
