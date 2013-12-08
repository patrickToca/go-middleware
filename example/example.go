package main

import (
	"github.com/shelakel/go-middleware"
	"log"
	"net/http"
	"time"
)

func logErrors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()
	log.Println("logErrors")
	next(w, r)
	log.Println("/logErrors")
}

func logRequests(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	started := time.Now()
	defer func() {
		elapsed := time.Since(started)
		log.Printf("%s: %s (%s)\n", r.Method, r.URL.RequestURI(), elapsed)
	}()
	log.Printf("logRequests: %s\n", r.URL.RequestURI())
	next(w, r)
	log.Printf("/logRequests: %s\n", r.URL.RequestURI())
}

func preFilter(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("preFilter")
	next(w, r)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
	<-time.After(2 * time.Millisecond)
}

func postFilter(w http.ResponseWriter, r *http.Request) {
	log.Println("postFilter")
}

func globalPostFilter(w http.ResponseWriter, r *http.Request) {
	log.Println("globalPostFilter")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", middleware.Compose(preFilter, helloWorld, postFilter))
	http.ListenAndServe(":8080", middleware.Compose(logErrors, logRequests, mux, globalPostFilter))
}
