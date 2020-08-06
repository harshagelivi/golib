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
	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	headers := lines[0]
	hmap := make(map[int]string, len(headers))
	for i, h := range headers {
		hmap[i] = h
	}
	for _, l := range lines[1:] {
		valmap := make(map[string]string, len(headers))
		for i, v := range l {
			valmap[hmap[i]] = v
		}
		cs.ch <- valmap
	}
	return nil
}
