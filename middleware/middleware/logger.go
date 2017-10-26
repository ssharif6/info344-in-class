package middleware

import "net/http"
import "time"
import "log"

//TODO: implement a Logger middleware handler
type Logger struct {
	Handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
}

// NewLogger creates a new instance of Logger middleware
func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler}
}
