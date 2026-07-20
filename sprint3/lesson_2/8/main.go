package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	data := []byte{12, 255, 129, 2, 1, 2, 255, 130, 0, 1, 12,
		0, 0, 17, 255, 130, 0, 2, 6, 72, 101, 108, 108,
		111, 44, 5, 119, 111, 114, 108, 100}
	var byfer *bytes.Buffer
	byfer = bytes.NewBuffer(data)
	array := make([]string, 0)
	dec := gob.NewDecoder(byfer)
	if err := dec.Decode(&array); err != nil {
		panic(err)
	}
	fmt.Print(array)
}
