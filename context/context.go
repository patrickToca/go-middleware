// Copyright 2013 Shelakel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package context allows for storing values on a *Request and is a middleware implementation of github.com/gorilla/context.
package context

import (
	"net/http"
	"sync"
)

const errContextNotInitialized string = "Use context.Initialize as middleware to scope the Request context before usage"

var (
	mutex sync.Mutex
	data  = make(map[*http.Request]map[interface{}]interface{})
)

// Scopes the Request context so that Set, Get, GetOk, Delete and Clear may be used.
func Initialize() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		mutex.Lock()
		if data[r] == nil {
			data[r] = make(map[interface{}]interface{})
		}
		mutex.Unlock()
		defer func() {
			mutex.Lock()
			delete(data, r)
			mutex.Unlock()
		}()

		next(w, r)
	}
}

// Set stores a value for a given key in a given request.
func Set(r *http.Request, key, val interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if data[r] == nil {
		panic(errContextNotInitialized)
	}
	data[r][key] = val
}

// Get returns a value stored for a given key in a given request.
func Get(r *http.Request, key interface{}) interface{} {
	mutex.Lock()
	defer mutex.Unlock()
	if data[r] == nil {
		panic(errContextNotInitialized)
	}
	return data[r][key]
}

// GetOk returns stored value and presence state like multi-value return of map access.
func GetOk(r *http.Request, key interface{}) (interface{}, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if data[r] == nil {
		panic(errContextNotInitialized)
	}
	val, ok := data[r][key]
	return val, ok
}

// Delete removes a value stored for a given key in a given request.
func Delete(r *http.Request, key interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if data[r] == nil {
		panic(errContextNotInitialized)
	}
	delete(data[r], key)
}

// Clear removes all values stored for a given request.
//
// This is usually called by a handler wrapper to clean up request
// variables at the end of a request lifetime. See ClearHandler().
func Clear(r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	if data[r] == nil {
		panic(errContextNotInitialized)
	}
	data[r] = make(map[interface{}]interface{})
}
