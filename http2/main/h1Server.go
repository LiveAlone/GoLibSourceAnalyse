package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handle)}

	log.Printf("Serving on https://0.0.0.0:8080")
	log.Fatal(srv.ListenAndServeTLS("data/server.crt", "data/server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	ct, err := w.Write([]byte("Hello World"))
	log.Printf("ct is %v, error %v", ct, err)
}
