package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []IServer{
		newServer("https://www.microsoft.com"),
		newServer("https://www.bing.com"),
		newServer("https://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)

	http.ListenAndServe(":"+lb.port, nil)
}
