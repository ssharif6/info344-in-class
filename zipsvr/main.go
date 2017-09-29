package main

import (
	"net/http"
	"log"
	"fmt"
	"runtime"
	"encoding/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloApi)
	mux.HandleFunc("/memory", memoryHandler)
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}

func helloApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello called")
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello, World " + name))
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats.HeapAlloc)
}


