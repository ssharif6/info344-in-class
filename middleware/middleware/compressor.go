package middleware

import "net/http"
import "strings"

import "compress/gzip"

//TODO: implement a gzip compression middleware
//handler that compresses the response stream
//using a gzip.Writer, if the client says it
//can handle that encoding. Check for a request
//header named Accept-Encoding, and if it's value
//contains the string "gzip", you can compress the
//response. Otherwise, don't compress the response

type GzipCompressor struct {
	handler http.Handler
}

func (c *GzipCompressor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if val := r.Header.Get("Accept-Encoding"); strings.Contains(val, "gzip") {
		gz := gzip.NewWriter(w)
		defer gz.Close()
	}
	return
}

// NewGzipCompressor creates new instance of Gzip middleware
func NewGzipCompressor(handler http.Handler) *GzipCompressor {
	return &GzipCompressor{handler}
}
