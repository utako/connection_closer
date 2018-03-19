package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Print("Starting...\n")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ConnState:      ConnStateListener,
	}

	panic(s.ListenAndServe())
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello")
}

func ConnStateListener(c net.Conn, cs http.ConnState) {
	if cs == http.StateNew {
		fmt.Println("New connection! Closing.")
		c.Close()
	} else {
		fmt.Printf("CONN STATE: %v, %v\n", cs, c)
	}
	fmt.Printf("CONN STATE: %v, %v\n", cs, c)
}
