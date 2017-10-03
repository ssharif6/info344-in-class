package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/ssharif6/info344-in-class/zipsvr/handlers"
	"github.com/ssharif6/info344-in-class/zipsvr/models"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		// Don't do log.Fatal for http handlers
		log.Fatal("Error loading zips: %v", err)
	}
	log.Printf("Loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}

	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloApi)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: "/zips/",
	}
	mux.Handle("/zips/", cityHandler)
	log.Fatal(http.ListenAndServe(addr, mux))
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
