package main

import (
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type()

	for i := 0; i < objType.NumField(); i++ {
		if v, ok := vobj.Field(i).Interface().(int); ok {
			limitTag, isLimit := objType.Field(i).Tag.Lookup("limit")
			if !isLimit {
				continue
			}
			limits := strings.Split(limitTag, ",")
			if v < Str2Int(limits[0]) {
				return false
			}
			if len(limits) > 1 && v > Str2Int(limits[1]) {
				return false
			}
		}
	}
	return true
}
