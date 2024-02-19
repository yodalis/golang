package main

import "net/http"

func main() {
	mux := http.NewServeMux() // Usado para ter mais controle das configurações do servidor\
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

// ServeHTTP implements http.Handler.
func (blog) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}
