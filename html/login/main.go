package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/", func(w http.ResponseWriter, r1 *http.Request) {
		r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
		http.ServeFile(w, r1, "./public/index.html")
	})
	http.ListenAndServe(":5000", nil)
	fmt.Println("1111")

}
