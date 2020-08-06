package work

import (
	"net/http"
)

type HttpSink struct {
	Url           string
	AuthRefresher AuthRefresher
	Client        *http.Client
}

func (hs *HttpSink) Do(p []interface{}) error {
	return nil
}
