package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewDefaultTransport ...
func NewDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// NewInsecureTransport ...
func NewInsecureTransport() *http.Transport {
	transport := NewDefaultTransport()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return transport
}
