package proxy

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type key int

const accessDenied key = 0
const accessAllowed key = 1

type directorError struct {
	err        error
	statusCode int
}

func NewDirector(target *url.URL) *Director {
	return &Director{
		TargetUrl: target,
	}
}

type Director struct {
	TargetUrl *url.URL
}

func (d *Director) Director(req *http.Request) {
	req.URL.Scheme = d.TargetUrl.Scheme
	req.URL.Host = d.TargetUrl.Host
	req.URL.Path = d.TargetUrl.Path

	if req.Method == http.MethodPost {
		err := errors.New("Not Allowed")
		*req = *req.WithContext(context.WithValue(req.Context(), accessDenied, &directorError{err: err, statusCode: http.StatusMethodNotAllowed}))
		return
	}
	*req = *req.WithContext(context.WithValue(req.Context(), accessAllowed, ""))
}

func (d *Director) RoundTrip(req *http.Request) (*http.Response, error) {
	if err, ok := req.Context().Value(accessDenied).(*directorError); ok && err != nil {
		return &http.Response{
			StatusCode: err.statusCode,
			Body:       ioutil.NopCloser(bytes.NewBufferString(err.err.Error())),
		}, nil
	}

	return http.DefaultTransport.RoundTrip(req)
}
