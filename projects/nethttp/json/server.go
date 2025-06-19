package nhj

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func StartServer(addr string, wg *sync.WaitGroup) (server *http.Server,
	url string) {
	url = "http://" + addr + "/echo"
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	bl := &bufferedListener{Listener: l, bufSize: common.IOBufSize}
	server = &http.Server{Handler: mux}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Serve(bl); err != http.ErrServerClosed {
			panic(fmt.Sprintf("Server error: %v\n", err))
		}
	}()
	time.Sleep(500 * time.Millisecond)
	return server, url
}

func CloseServer(server *http.Server, wg *sync.WaitGroup) (err error) {
	err = server.Close()
	if err != nil {
		return
	}
	wg.Wait()
	return
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data common.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	time.Sleep(common.Delay)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

type bufferedListener struct {
	net.Listener
	bufSize int
}

func (bl *bufferedListener) Accept() (net.Conn, error) {
	conn, err := bl.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return &bufferedConn{
		Conn: conn,
		r:    bufio.NewReaderSize(conn, bl.bufSize),
		w:    bufio.NewWriterSize(conn, bl.bufSize),
	}, nil
}
