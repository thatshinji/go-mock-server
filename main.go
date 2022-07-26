package main

import (
	"encoding/json"
	"go-mock-server/src"
	"log"
	"net/http"
	"os"
)

var data map[string]interface{}

func returnResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	path := r.URL.Path
	if d := data[path]; d != nil {
		stringData, err := json.Marshal(data[path])
		if err != nil {
			log.Fatal(err)
		}
		w.Write(stringData)
	}
}

var port string = "3000"

func main() {
	if l := len(os.Args); l > 1 {
		port = os.Args[1]
	}
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
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("listenAndServer fail")
	}
}
