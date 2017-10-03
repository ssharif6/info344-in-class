package handlers

import "github.com/ssharif6/info344-in-class/zipsvr/models"
import "net/http"
import "strings"
import "encoding/json"

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// receiver function that allows you to attach objects in the http handler!!
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := strings.ToLower(r.URL.Path[len(ch.PathPrefix):])
	if len(cityName) == 0 {
		// NOTE: USE THIS FOR HTTP ERRORS
		http.Error(w, "Please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
