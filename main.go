package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", good_app)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func good_app(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "good app.")
}
