package nhj

import (
	"bufio"
	"context"
	"net"
	"net/http"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func MakeBufferedClient(connsCount int) *http.Client {
	transport := &http.Transport{
		MaxConnsPerHost:     connsCount,
		MaxIdleConnsPerHost: connsCount,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
			// Wrap with buffered read/write
			return &bufferedConn{
				r:    bufio.NewReaderSize(conn, common.IOBufSize),
				w:    bufio.NewWriterSize(conn, common.IOBufSize),
				Conn: conn,
			}, nil
		},
	}

	return &http.Client{Transport: transport}
}
