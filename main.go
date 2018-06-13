package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/taland/reverse-proxy/proxy"
)

var (
	addr     = flag.String("addr", ":5555", "TCP address to listen to")
	destHost = flag.String("desthost", "http://localhost:5000", "Destination host")
)

func main() {
	flag.Parse()

	u, err := url.Parse(*destHost)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", proxy.NewProxy(u).Handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
