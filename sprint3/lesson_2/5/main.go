package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Slice []byte

func (s Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(s))
}

func (s *Slice) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v, err := hex.DecodeString(tmp)
	if err != nil {
		return err
	}
	*s = v
	return nil
}

type MySlice struct {
	ID    int
	Slice Slice
}

func main() {
	ret, err := json.Marshal(MySlice{ID: 7, Slice: []byte{1, 2, 3, 10, 11, 255}})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
	var result MySlice
	if err = json.Unmarshal(ret, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
