package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Printf("%s", dump)
		w.Write(dump)
		return
	})
	http.ListenAndServe(":8080", nil)
}
