package work

import (
	"sync/atomic"
	"time"
)

type AuthRefresher struct {
	Token atomic.Value
	Refresher func()string
	AuthHeader string
}

func (ar *AuthRefresher) GetToken() string {
	return ar.Token.Load().(string)
}

func (ar *AuthRefresher) Start(secs int)  {
	for {
		select {
		case <-time.Tick(time.Second * time.Duration(secs)):
			ar.Token.Store(ar.Refresher())
		}
	}
}
