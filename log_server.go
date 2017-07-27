package main

import (
	"log"
	"net/http"
	"time"
)

type logServer struct {
	Next   http.Handler
	Logger *log.Logger
}

type logResponseWriter struct {
	http.ResponseWriter
	Code int
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

func (s *logServer) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	lw := &logResponseWriter{w, 200}
	defer func(b time.Time) {
		log.Printf("%s %s %d %v", request.Method, request.URL.String(), lw.Code, time.Since(b))
	}(time.Now())
	s.Next.ServeHTTP(lw, request)
}
