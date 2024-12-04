package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	h := http.NewServeMux()
	h.HandleFunc("POST /hello", helloHendler)
	newH := Middleware(h)
	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: newH,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func helloHendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(now)
		log.Printf("[%s] %q %v\n", r.Method, r.URL.Path, elapsed)
	})
}
