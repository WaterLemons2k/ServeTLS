package main

import (
	"crypto/tls"
	_ "embed"
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":443", "Address to listen on")

	//go:embed server.crt
	serverCrt []byte

	//go:embed server.key
	serverKey []byte
)

func main() {
	flag.Parse()

	serveTLS(serverCrt, serverKey)
}

func serveTLS(certPEM, keyPEM []byte) {
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatalf("Failed to parse certificate: %v", err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	srv := &http.Server{
		Addr:      *addr,
		TLSConfig: cfg,
	}

	log.Println("Listening on", *addr)
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
