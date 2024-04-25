package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cansel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cansel()

	mux := http.NewServeMux()
	mux.HandleFunc("/", hendler)

	srv := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}
	go startServer(srv)
	<-ctx.Done()
	srv.Shutdown(context.Background())
	fmt.Println("Server stoped gracefully")
}

func hendler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))

}

func startServer(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Server closed with error", err)
	}

}
