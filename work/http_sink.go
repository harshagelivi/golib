package work

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpSink struct {
	Url    string
	Method string
	Auth   *AuthRefresher
	//ClientPool    ClientPool
	Client    *http.Client
	BodyMaker func([]interface{}) io.Reader
}

func (hs *HttpSink) Do(p []interface{}) error {
	req, err := http.NewRequest(hs.Method, hs.Url, hs.BodyMaker(p))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", hs.Auth.AuthHeader)
	req.Header.Add("X-Auth-Token", hs.Auth.GetToken())
	//client := hs.ClientPool.GetClient()
	start := time.Now()
	res, err := hs.Client.Do(req)
	if time.Now().Unix()%10 == 0 {
		fmt.Printf("http req took %v\n", time.Since(start))
	}
	if err != nil {
		return err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("Error while closing %v", err)
		}
	}()
	if res.StatusCode != 200 {
		fmt.Println(res)
	}
	return nil
}
