package main

import (
	"fmt"
	"net/http"
	"strings"
)

type handler http.HandlerFunc

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-tpe", "text/plain")
	fmt.Fprintln(w, "Hello canonical world")
}

func wrapCanonicalHost(f handler, chost string) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		hostPort := strings.Split(r.Host, ":")
		host := hostPort[0]
		if host != chost {
			fmt.Println("redirect to", chost)
			hostPort[0] = chost
			url := "http://" +
				strings.Join(hostPort, ":") +
				r.URL.String()

			http.Redirect(w, r, url, 301)
			return
		}
		f(w, r)
	}
}

func main() {
	handler := wrapCanonicalHost(hello, "localhost")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
