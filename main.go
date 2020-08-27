package main

import (
	"fmt"
	"net/http"
)

func routers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Github WebHook commit 3")
	})

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HOla Probando commit 3")
	})

}

func main() {
	routers()
	http.ListenAndServe(":4000", nil)
}
