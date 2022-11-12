package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = 8080
const serverKeyPath = "keys/server.key"
const serverCertPath = "keys/server.pem"

func main() {
	// 添加监听器
	srv := &http.Server{Addr: fmt.Sprintf(":%d", serverPort), Handler: http.HandlerFunc(handle)}

	// 监听启动SSL
	log.Println("server start to listener https://localhost:8080/")
	log.Fatal(srv.ListenAndServeTLS(serverCertPath, serverKeyPath))

	// 非SSL监听
	//log.Println("server start to listener http://localhost:8080/")
	//log.Fatal(srv.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
	// http 协议
	log.Printf("Got connection: %s", r.Proto)

	ct, err := w.Write([]byte("Hello World"))

	log.Printf("ct is %v, error %v", ct, err)
}
