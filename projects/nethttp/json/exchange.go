package nhj

import (
	"bytes"
	"encoding/json"
	"io"
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

func send(url string, data common.Data, client *http.Client) (common.Data, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return common.Data{}, err
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return common.Data{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.Data{}, err
	}
	var adata common.Data
	if err := json.Unmarshal(body, &adata); err != nil {
		return common.Data{}, err
	}
	return adata, nil
}
