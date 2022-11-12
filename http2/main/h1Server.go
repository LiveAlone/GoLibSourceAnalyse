package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handle)}
	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	//log.Fatal(srv.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	ct, err := w.Write([]byte("Hello World"))
	log.Printf("ct is %v, error %v", ct, err)
}
