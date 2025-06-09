package nhj

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ymz-ncnk/go-client-server-communication-benchmarks/common"
)

func StartServer(addr string, wg *sync.WaitGroup) (server *http.Server, url string) {
	url = "http://" + addr + "/echo"
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)

	server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
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
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var data common.Data
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
