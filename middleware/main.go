package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ssharif6/info344-in-class/middleware/middleware"
	"github.com/ssharif6/info344-in-class/middleware/handlers"
	"fmt"
)

const defaultAddr = ":4000"

func main() {
	addr := os.Getenv("ADDR")
	fmt.Println(addr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	mux := http.NewServeMux()
	// mux.HandleFunc("/hello", handlers.HelloHandler)

	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/time", handlers.TimeHandler)

	//TODO: wrap the mux with the Logger middleware and gzip
	wrappedMux := middleware.NewLogger(mux)

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
