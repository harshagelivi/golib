package work

import (
	"net/http"
	"time"
)

type ClientPool struct {
	clients []*http.Client
	Size int64

}

func (cp *ClientPool) Init(clientFactory func()(*http.Client, error)) error {
	cp.clients = make([]*http.Client, cp.Size)
	for i := int64(0); i < cp.Size; i++ {
		c, err := clientFactory()
		if err != nil {
			return err
		}
		cp.clients[i] = c
	}
	return nil
}

func (cp *ClientPool) GetClient() *http.Client {
	return cp.clients[time.Now().Unix()%cp.Size]
}
