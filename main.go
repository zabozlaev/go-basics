package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type requestBody struct {
	Message string `json:"msg"`
}

func writeCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Add("Access-Controll-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, content-type, Content-Length, Accept-Encoding")
}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	writeCorsHeaders(&w)

	method := req.Method

	contentType := req.Header.Get("Content-Type")

	if method != "POST" && method != "OPTIONS" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("MethodNotAllowed"))
		return
	}

	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid content type"))
		return
	}

	var body requestBody

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid request body"))
		return
	}

	fmt.Printf("Received message: %s \n", body.Message)

	w.WriteHeader(http.StatusAccepted)

	w.Write([]byte("Successfuly received your message!"))
}

func main() {
	http.HandleFunc("/", requestHandler)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		log.Fatal(err)
	}
}
