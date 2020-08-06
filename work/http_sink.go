package work

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpSink struct {
	url string
}

func (hs *HttpSink) Do(p []interface{}) error {
	byts, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = http.Post(hs.url, "application/json", bytes.NewReader(byts))
	return err
}
