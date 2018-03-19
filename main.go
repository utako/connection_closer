package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", internal_server_error)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func internal_server_error(res http.ResponseWriter, req *http.Request) {
	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
