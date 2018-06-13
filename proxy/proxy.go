package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewProxy(target *url.URL) *Proxy {
	d := NewDirector(target)

	rp := &httputil.ReverseProxy{
		Director:  d.Director,
		Transport: d,
	}

	return &Proxy{proxy: rp}
}

type Proxy struct {
	proxy *httputil.ReverseProxy
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-PROXY", "v1")
	p.proxy.ServeHTTP(w, r)
}
