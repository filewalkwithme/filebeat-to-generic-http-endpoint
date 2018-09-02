package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}
		fmt.Printf("INFO: %s\n", string(buf))
	})

	http.ListenAndServe(":8080", nil)
}
