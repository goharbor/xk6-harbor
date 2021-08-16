package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

// NewDefaultTransport ...
func NewDefaultTransport() *http.Transport {
	forceAttemptHTTP2 := cast.ToBool(GetEnv("HTTP_FORCE_ATTEMPT_HTTP2", "true"))
	disableKeepAlives := cast.ToBool(GetEnv("HTTP_DISABLE_KEEP_ALIVES", "false"))
	maxIdleConns := cast.ToInt(GetEnv("HTTP_MAX_IDLE_CONNS", "100"))
	maxConnsPerHost := cast.ToInt(GetEnv("HTTP_MAX_CONNS_PER_HOST", "0"))
	maxIdleConnsPerHost := cast.ToInt(GetEnv("HTTP_MAX_IDLE_CONNS_PER_HOST", "0"))

	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		DisableKeepAlives:     disableKeepAlives,
		ForceAttemptHTTP2:     forceAttemptHTTP2,
		MaxIdleConns:          maxIdleConns,
		MaxConnsPerHost:       maxConnsPerHost,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
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
