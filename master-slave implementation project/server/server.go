package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	// "io/ioutil"

)

func main() {
	// Create a slice of URLs for the slaves
	urls := []*url.URL{
		{Scheme: "http", Host: "localhost:8080"},
		{Scheme: "http", Host: "localhost:8081"},
		{Scheme: "http", Host: "localhost:8082"},
		{Scheme: "http", Host: "localhost:8083"},
	}

	// Create a reverse proxy that load balances requests across the slaves
	director := func(req *http.Request) {
		target := urls[rand.Int()%len(urls)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}

	// Handle incoming requests by load balancing them across the slaves
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		
		fmt.Printf("server received the client request.")
		proxy.ServeHTTP(w, r)
		
		
		// call slave0 to combine the parts
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

	})
	fmt.Println()
	fmt.Println("master server working and wait client request..!")
	fmt.Println()
	log.Fatal(http.ListenAndServe(":9090", nil))
}
