// Copyright 2013 Shelakel. All rights reserved.
// Use of this source code is governed by a BSD-style license
// license that can be found in the LICENSE file.
// github.com/shelakel/go-middleware

// Package middleware provides support for composing on demand and reusable filters/middleware
//
// See github.com/shelakel/go-middleware/example for an example on how
// Compose is used to create middleware and pre/post filters.
package middleware

import (
	"log"
	"net/http"
	"reflect"
)

// The MiddlewareFunc type is an adapter to allow the use of ordinary functions as Middleware.
// If f is a function with the appropriate signature, MiddlewareFunc(f) is a Middleware object that calls f.
type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// Process calls f(w, r, next).
func (fn MiddlewareFunc) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fn(w, r, next)
}

// Middleware allows for creating reusable middleware that require initial configuration
type Middleware interface {
	Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

func nullHandlerFunc(w http.ResponseWriter, r *http.Request) {}

// Compose chains together the middleware and returns a http.HandlerFunc (which is compatible with http.Handler).
// Allowed middleware variants include:
// func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request)),
// func(http.ResponseWriter, *http.Request, http.HandlerFunc),
// func(http.ResponseWriter, *http.Request, http.Handler),
// Middleware, MiddlewareFunc,
// func(http.ResponseWriter, *http.Request),
// http.HandlerFunc, http.Handler and func(http.Handler) http.Handler.
func Compose(middleware ...interface{}) http.HandlerFunc {
	var fn http.HandlerFunc = nullHandlerFunc

	for i := len(middleware) - 1; i >= 0; i-- {
		next := fn // Store variable locally for implicit capture by closure

		switch current := middleware[i].(type) {
		// MiddlewareFunc signatures
		case func(http.ResponseWriter, *http.Request, func(http.ResponseWriter, *http.Request)):
			fn = func(w http.ResponseWriter, r *http.Request) { current(w, r, next) }
		case func(http.ResponseWriter, *http.Request, http.HandlerFunc):
			fn = func(w http.ResponseWriter, r *http.Request) { current(w, r, next) }
		case func(http.ResponseWriter, *http.Request, http.Handler):
			fn = func(w http.ResponseWriter, r *http.Request) { current(w, r, next) }
		case Middleware: // MiddlewareFunc implements Middleware
			fn = func(w http.ResponseWriter, r *http.Request) { current.Process(w, r, next) }
		// http.HandlerFunc signatures
		case func(http.ResponseWriter, *http.Request):
			fn = func(w http.ResponseWriter, r *http.Request) { current(w, r); next(w, r) }
		case http.Handler: // http.HandlerFunc implements http.Handler
			fn = func(w http.ResponseWriter, r *http.Request) { current.ServeHTTP(w, r); next(w, r) }
		case func(http.Handler) http.Handler:
			fn = func(w http.ResponseWriter, r *http.Request) { current(next).ServeHTTP(w, r) }
		default:
			log.Panicf("Unsupported middleware type '%v' at index %d", reflect.TypeOf(current), i)
		}
	}

	return fn
}
