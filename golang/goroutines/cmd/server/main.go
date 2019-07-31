package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	listener = flag.String("listener", "127.0.0.1:1337", "The IP and port to listen on")
)

type waitResponse struct {
	Status string `json:"status"`
}

func waitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received connection from %s\n", r.RemoteAddr)
	time.Sleep(time.Second * 2)

	ret := waitResponse{
		Status: "done",
	}

	json.NewEncoder(w).Encode(ret)
}

func main() {
	flag.Parse()

	http.HandleFunc("/wait", waitHandler)
	http.ListenAndServe(*listener, nil)
}
