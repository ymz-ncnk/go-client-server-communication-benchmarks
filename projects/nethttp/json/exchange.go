package nhj

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func ExchangeQPS(url string, data common.Data, client *http.Client,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	adata, err := send(url, data, client)
	if err != nil {
		b.Error(err)
		return
	}
	if !common.EqualData(data, adata) {
		b.Error("unexpected result")
	}
}

func ExchangeFixed(url string, data common.Data, client *http.Client,
	copsD chan<- time.Duration,
	wg *sync.WaitGroup,
	b *testing.B,
) {
	defer wg.Done()
	start := time.Now()
	adata, err := send(url, data, client)
	if err != nil {
		b.Error(err)
		return
	}
	common.QueueCopD(copsD, time.Since(start))
	if !common.EqualData(data, adata) {
		b.Error("unexpected result")
	}
}

func send(url string, data common.Data, client *http.Client) (result common.Data,
	err error) {
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(data)
	if err != nil {
		return
	}
	resp, err := client.Post(url, "application/json", buf)
	if err != nil {
		return common.Data{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
