package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
1. Generate cert.pem and key.pem using this:

openssl req -x509 -nodes -days 365 -subj '/C=UK/ST=Example/L=Example/CN=localhost' -newkey rsa:1024 -keyout key.pem -out cert.pem

2. Run web-static-tls.go in the same directory

go run web-static-tls.go

3. Navigate to https://localhost:8443/about

*/
func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page over HTTPS")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The index page over HTTPS")
}
