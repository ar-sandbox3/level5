package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello/handler1", basicAuth(http.StripPrefix("/hello", http.HandlerFunc(handler1))))
	mux.Handle("/hello/handler2", basicAuth(http.StripPrefix("/hello", http.HandlerFunc(handler2))))
	http.ListenAndServe(":8080", mux)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func basicAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		_, _, ok := req.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", "Basic real=hsi")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Processing request.
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}
