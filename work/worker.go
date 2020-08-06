package work

import (
	"fmt"
	"time"
)

type Worker struct {
	DataChannel <-chan interface{}
	DstSink     Sink
	NumWorkers  int
	BuffSize    int
}

func (w *Worker) Run() {
	for i := 0; i < w.NumWorkers; i++ {
		go w.sessionRun()
	}
}

func (w *Worker) sessionRun() {
	sessionData := make([]interface{}, w.BuffSize)
	sessionIdx := 0
	fmt.Printf("Starting at %v", time.Now())
	for datum := range w.DataChannel {
		sessionData[sessionIdx] = datum
		sessionIdx++
		if sessionIdx == w.BuffSize {
			e := w.DstSink(sessionData)
			sessionIdx = 0
			if e != nil {
				// todo add more handling/exponential retry
				fmt.Println(e)
			}
		}
	}
	fmt.Printf("Ending at %v", time.Now())
}
