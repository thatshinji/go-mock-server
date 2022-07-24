package main

import (
	"fmt"
	"go-mock-server/src"
	"log"
	"net/http"
)

var data map[string]interface{}

func returnResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	path := r.URL.Path
	//w.Write(data[path])
	if d := data[path]; d != nil {
		fmt.Println(d, "data")
	}
}

func main() {
	path, err := src.ParsePath()
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := src.ReadJSON(path)
	if err != nil {
		log.Fatal(err)
	}
	m, err := src.DecodeJSONString(bytes)
	data = m
	if err != nil {
		log.Fatal("decode fail: ", err)
	}
	for addr, _ := range data {
		http.HandleFunc(addr, returnResponse)
	}
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("listenAndServer fail")
	}
}
