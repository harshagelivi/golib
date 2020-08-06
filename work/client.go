package work

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func NewClient(timeoutSecs int, maxIdle int) *http.Client {
	timeout := time.Second * time.Duration(timeoutSecs)
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
				DualStack: true,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          maxIdle,
			MaxIdleConnsPerHost:   maxIdle,
			MaxConnsPerHost:       maxIdle,
			IdleConnTimeout:       timeout * 2,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: timeout,
	}
}
