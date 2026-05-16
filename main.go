package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	name := os.Getenv("NAME")
	if name == "" {
		name = fmt.Sprintf("backend-%s", port)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, name)
		w.Header().Set("X-Backend", name)
		fmt.Fprintf(w, "Hello from %s\n", name)
	})

	mux.HandleFunc("/api/auth/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, name)
		w.Header().Set("X-Backend", name)
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/auth/login":
			fmt.Fprint(w, `{"endpoint":"login","status":"ok"}`)
		case "/api/auth/register":
			fmt.Fprint(w, `{"endpoint":"register","status":"ok"}`)
		case "/api/auth/logout":
			fmt.Fprint(w, `{"endpoint":"logout","status":"ok"}`)
		default:
			fmt.Fprint(w, `{"endpoint":"auth","status":"ok"}`)
		}
	})

	mux.HandleFunc("/api/buy", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, name)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Backend", name)
		fmt.Fprint(w, `{"endpoint":"buy","status":"ok"}`)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, name)
		w.Header().Set("X-Backend", name)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"healthy"}`)
	})

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("%s listening on %s", name, addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
