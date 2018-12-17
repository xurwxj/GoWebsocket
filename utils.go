package gowebsocket

import (
	"net/http"
	"net/url"
)

func BuildProxy(Url string) func(*http.Request) (*url.URL, error) {
	if uProxy, err := url.Parse(Url); err != nil {
		// log.Fatal("Error while parsing url ", err)
		return nil
	} else {
		return http.ProxyURL(uProxy)
	}
}
