package work

import (
	"encoding/csv"
	"os"
)

type CsvSrc struct {
	fileName string
	ch chan<- interface{}
}

func (cs *CsvSrc) Do() error {
	f, err := os.Open(cs.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	for _, l := range lines {
		cs.ch <- l
	}
	return nil
}
