package main

import "net/http"

// StripPrefix is an example adapter pattern for adapting middleware of signature func(...options, http.Handler) htt.Handler
// ex. http.StripPrefix(prefix, handler) to a MiddlewareFunc: func(http.ResponseWriter, *http.Request, http.HandlerFunc)
// Example adapter pattern originally provided to address concerns as expressed at http://www.reddit.com/r/golang/comments/1sdoji/pure_go_middleware/
func StripPrefix(prefix string) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		http.StripPrefix(prefix, next).ServeHTTP(w, r)
	}
}
