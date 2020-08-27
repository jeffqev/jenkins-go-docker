package main

import (
	"fmt"
	"net/http"
)

func routers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Github WebHook commit 2")
	})

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HOla Probando")
	})

}

func main() {
	routers()
	http.ListenAndServe(":4000", nil)
}
