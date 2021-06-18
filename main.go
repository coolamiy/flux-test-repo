package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello ")
}
