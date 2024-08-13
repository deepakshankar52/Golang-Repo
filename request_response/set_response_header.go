package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8085", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	w.Header().Add("foo", "bar1")
	w.Header().Add("foo", "bar2")

	resp := make(map[string]string)
	resp["message"] = "success"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}