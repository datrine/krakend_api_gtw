package middleware

import (
	"net/http"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request) bool
type MiddlewareHandlerFunc func(HandlerFunc) func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

/*
type MiddlewareFuncDef func(f HandlerFunc) func(w http.ResponseWriter, req *http.Request)

type Midd func(h HandlerFunc) HandlerFunc

var MiddlewareFunc MiddlewareFuncDef
*/
func Chain(arr ...MiddlewareFunc) MiddlewareHandlerFunc {
	return func(h HandlerFunc) func(w http.ResponseWriter, req *http.Request) {
		return func(w http.ResponseWriter, req *http.Request) {
			var val bool = true
			for _, mid := range arr {
				if val && req != nil {
					val = mid(w, req)
				} else {
					return
				}
			}
			if val {
				h(w, req)
			}
		}
	}
}
