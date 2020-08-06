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
	start := time.Now()
	minuteStart := time.Now()
	numRecs := 0
	totalRecs := 0
	loop:
	for {
		select {
		case datum, ok := <-w.DataChannel:
			if !ok {
				break loop
			}
			sessionData[sessionIdx] = datum
			sessionIdx++
			numRecs++
			if sessionIdx == w.BuffSize {
				e := w.DstSink(sessionData)
				sessionIdx = 0
				if e != nil {
					// todo add more handling/exponential retry
					fmt.Println(e)
				}
			}
		case <-time.Tick(time.Second * 60):
			fmt.Printf("Handled %d in %v slot\n", numRecs, time.Since(minuteStart))
			minuteStart = time.Now()
			totalRecs += numRecs
			numRecs = 0
		}
	}
	fmt.Printf("Handled %d in %v\n", totalRecs, time.Since(start))
}
