package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Got error while reading body: %v", err)))
			return
		}
		w.Write([]byte(fmt.Sprintf("Body length: %d Body: %q", len(b), b)))
	})

	http.ListenAndServe(":80", nil)
}
