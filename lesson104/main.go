package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	work(ctx)
	select {
	case <-ctx.Done():
		fmt.Fprintln(w, "not done")
		return
	default:
	}
	fmt.Fprintln(w, "done")
}

func work(ctx context.Context) {
	for i := 0; i < 3_000_000; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		fmt.Println(i)
	}
}
