package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	crt = `server.crt`
	key = `server.key`
)

func main() {
	// This goroutine is optional, but it's probably a decent idea to redirect all traffic to be over https.
	// If you don't serve anything on port 80 then your website just won't work when people try to use http.
	go func() {
		log.Println("Redirecting traffic from port 80 to 443 (http -> https)")
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(httpsRedirect)))
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhello)
	log.Println("Serving on port 443 (https)")
	log.Fatal(http.ListenAndServeTLS(":443", crt, key, mux))
}

// httpsRedirect is based on https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
func httpsRedirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
