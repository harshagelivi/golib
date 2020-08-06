package work

import (
	"fmt"
)

type PrintSink struct {
}

func (ps *PrintSink) Do(p []interface{}) error {
	fmt.Println(p)
	return nil
}
