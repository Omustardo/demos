package main

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/acme/autocert"
	"crypto/tls"
	"log"
)

func main() {
	go func() {
		log.Println("Redirecting traffic from port 80 to 443 (http -> https)")
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(httpsRedirect)))
	}()

	// TODO: Switch to autocert. https://godoc.org/golang.org/x/crypto/acme/autocert
	// It's currently broken. LetsEncrypt found a vulnerability with tls-sni-01 (a way of verifying that you own a domain)
	// and so they disabled it. This package is supposed to also work with tls-sni-02 but when I run it, it only gets
	// challenge types dns-01 and http-01 which aren't supported.
	m := autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.omustardo.com", "omustardo.com"),
		Email:      "omustardo@omustardo.com",
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}
	http.HandleFunc("/", sayhello)
	log.Fatal(s.ListenAndServeTLS("", ""))
}

func httpsRedirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello jacob!") // send data to client side
}
