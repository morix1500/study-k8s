package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hoge", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hoge"))
	})
	http.HandleFunc("/fuga", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Fuga"))
	})
	http.HandleFunc("/piyo", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("piyo"))
	})
	http.ListenAndServe(":3000", nil)
}
