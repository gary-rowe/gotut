package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func main() {

	// Start web server
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	// Create a mock page
	page := NewsAggPage{Title: "Amazing news aggregator", News: "some news"}

	// Parse the template file
	t, _ := template.ParseFiles("basic-templating.html")

	// Use Execute err to uncover problems in the template page
	fmt.Println(t.Execute(w, page))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Yo! I made a server</h1>")
}
