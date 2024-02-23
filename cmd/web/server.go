package main

import (
	"net/http"

	"github.com/ar-sandbox3/level5/authenticator"
)

func main() {
	mux := http.NewServeMux()
	a, err := authenticator.NewAuthenticator("https://sandbox.hsi.id")
	if err != nil {
		panic(err)
	}

	mux.Handle("/hello", auth(a, http.HandlerFunc(handler1)))

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

func auth(a *authenticator.Authenticator, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := []byte(authHeader)
		if err := a.Validate(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

func basicAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		_, _, ok := req.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", "Basic realm=hsi")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Processing request.
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}
