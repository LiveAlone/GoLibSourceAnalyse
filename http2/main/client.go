package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const url = "https://localhost:8080"
const clientCertPath = "keys/server.pem"

// params -version 2
var httpVersion = flag.Int("version", 1, "HTTP version")

func main() {
	flag.Parse()
	client := &http.Client{}

	// init transport
	tlsConfig, err := initCA()
	if err != nil {
		log.Printf("ca load error, cause:%v", err)
		return
	}

	// init 不同client http
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	case 2:
		client.Transport = &http2.Transport{TLSClientConfig: tlsConfig}
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf("Got response %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}

// Load CA
func initCA() (*tls.Config, error) {
	caCert, err := os.ReadFile(clientCertPath)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		ServerName: "www.p-pp.cn",
		RootCAs:    caCertPool,
	}, nil
}
