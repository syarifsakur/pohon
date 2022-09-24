package main

import (
	"fmt"
	"net/http"
)

type pohon struct {
	name  string  `json:"name"`
	price float64 `json:"price"`
}

type response struct {
	code    int         `json:"code"`
	message string      `json:"message"`
	data    interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/api/v1/trees", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	port := "8000"
	fmt.Println("server run on port",port)
	http.ListenAndServe(":"+port,nil)
}