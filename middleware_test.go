// Copyright 2013 Shelakel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// github.com/shelakel/go-middleware

package middleware

import (
	"net/http"
	"testing"
)

func TestComposeChainExecute(t *testing.T) {
	i := 1

	incrHandler := func(w http.ResponseWriter, r *http.Request) { i = i * 2 }
	incrMiddleware := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) { i = i * 3; next(w, r) }

	// Simple test to verify that all methods in a chain execute under standard conditions
	Compose(incrHandler, incrMiddleware, incrHandler, incrMiddleware)(nil, nil)
	if i != 36 {
		t.Fatalf("Expected i to be 36, instead it was %d", i)
	}
}

type stubHandler struct{}

func (s *stubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func stubMiddleware1(w http.ResponseWriter, r *http.Request, next func(http.ResponseWriter, *http.Request)) {
	next(w, r)
}

func stubMiddleware2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) { next(w, r) }

func stubMiddleware3(w http.ResponseWriter, r *http.Request, next http.Handler) { next.ServeHTTP(w, r) }

type stubMiddleware struct{}

func (s *stubMiddleware) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func TestComposeMiddlewareVariants(t *testing.T) {
	fn1 := nullHandlerFunc
	fn2 := http.HandlerFunc(nullHandlerFunc)
	fn3 := http.Handler(&stubHandler{})
	fn4 := stubMiddleware1
	fn5 := stubMiddleware2
	fn6 := stubMiddleware3
	fn7 := Middleware(&stubMiddleware{})
	fn8 := MiddlewareFunc(stubMiddleware2)
	fn9 := func(h http.Handler) http.Handler { return h }

	Compose(fn1, fn2, fn3, fn4, fn5, fn6, fn7, fn8, fn9)
}

func TestComposeNilChain(t *testing.T) {
	fn := Compose()

	if fn == nil {
		t.Fatal("Expected fn to not be nil")
	}
}

func TestComposeInvalidMiddleware(t *testing.T) {

}
