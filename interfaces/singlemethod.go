package interfaces

import "net/http"

type handlerFunc func(http.ResponseWriter, *http.Request)

// function that calls itself
func (f handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
