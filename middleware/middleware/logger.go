package middleware

import "net/http"
import "time"
import "log"

//TODO: implement a Logger middleware handler
type Logger struct {
	Handler http.Handler
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lrw := &loggingResponseWriter{w, http.StatusOK}
	start := time.Now()
	l.Handler.ServeHTTP(lrw, r)
	log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
}

// NewLogger creates a new instance of Logger middleware
func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler}
}
