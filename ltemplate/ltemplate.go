package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"net/http"
	"reflect"
)

func getBigger[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

type Apple[T cmp.Ordered] struct{}

func (Apple[T]) getBigger(a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

type GetUserRequest struct{}
type GetBookRequest struct{}

func httpRPC[T GetUserRequest | GetBookRequest](request T) {
	url := "http://127.0.0.1/"
	tp := reflect.TypeOf(request)
	switch tp.Name() {
	case "GetUserRequest":
		url = url + "user"
	case "GetBookRequest":
		url = url + "book"
	default:
		panic("unknown type")
	}
	bs, _ := json.Marshal(request)
	http.Post(url, "application/json", bytes.NewBuffer(bs))
}

func main() {

	//bigger := getBigger[int32](1, 2)
	//fmt.Printf("%d \n", bigger)
	a := Apple[int]{}
	a.getBigger(1, 2)
}
