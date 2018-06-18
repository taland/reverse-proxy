package proxy

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQueryString(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer backend.Close()

	u, _ := url.Parse(backend.URL)
	d := NewDirector(u)

	proxy := httptest.NewServer(&httputil.ReverseProxy{Director: d.Director, Transport: d})
	defer proxy.Close()

	for _, params := range []struct {
		url  string
		code int
	}{
		{
			url:  "/something",
			code: 200,
		},
		{
			url:  "/something?p0=&p1=foo",
			code: 200,
		},
	} {
		req, _ := http.NewRequest("GET", proxy.URL+params.url, nil)
		res, _ := http.DefaultClient.Do(req)

		assert.Equal(t, params.code, res.StatusCode)
	}
}
