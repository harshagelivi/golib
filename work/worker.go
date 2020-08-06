package work

import "fmt"

type Worker struct {
	dataChannel chan<- interface{}
	dstSink Sink
	numWorkers int
	buffSize int
}

func (w *Worker) Run()  {
	for i := 0; i < w.numWorkers; i++ {
		go w.sessionRun()
	}
}

func (w *Worker) sessionRun()  {
	sessionData := make([]interface{}, w.buffSize)
	sessionIdx := 0
	for datum := range w.dataChannel {
		sessionData[sessionIdx] = datum
		sessionIdx++
		if sessionIdx == w.buffSize {
			e := w.dstSink(sessionData)
			sessionIdx = 0
			if e != nil {
				// todo add more handling/exponential retry
				fmt.Println(e)
			}
		}
	}
}
