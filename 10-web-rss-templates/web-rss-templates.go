package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*

Blends the RSS feed reader and the web server components together

*/

// The site map "master" struct <sitemap>...<loc>
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// The news "slave" sitemap has a more detailed XML schema
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// The /agg page handler
func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	// GET from upstream feed
	// Necessary to have headers set to avoid tar pit from upstream
	sitemapReq, _ := http.NewRequest("GET", "https://www.washingtonpost.com/news-sitemaps/index.xml", nil)
	sitemapReq.Header.Set("Connection", "Keep-Alive")
	sitemapReq.Header.Set("Accept-Language", "en-US")
	sitemapReq.Header.Set("User-Agent", "Mozilla/5.0")

	sitemapResp, _ := client.Do(sitemapReq)
	sitemapBytes, _ := ioutil.ReadAll(sitemapResp.Body)
	sitemapResp.Body.Close()

	// Parse the body as XML using location type
	var sitemapIndex SitemapIndex
	xml.Unmarshal(sitemapBytes, &sitemapIndex)

	// Create a map to store the news aggregator page
	newsItems := make(map[string]News)

	// Iterate over sitemap locations using range (index and element)
	// and put the results in a map
	for _, location := range sitemapIndex.Locations {

		// Location may have whitespace
		location = strings.TrimSpace(location)

		// Only interested in politics and opinions to keep traffic down
		if strings.Contains(location, "politics.xml") ||
			strings.Contains(location, "opinions.xml") {

			// Print the URL
			fmt.Printf("GET: %s\n", location)

			// GET from upstream feed
			// Necessary to have headers set to avoid tar pit from upstream
			newsReq, err := http.NewRequest("GET", location, nil)
			if err != nil {
				log.Fatal(err)
			}
			newsReq.Header.Set("Connection", "Keep-Alive")
			newsReq.Header.Set("Accept-Language", "en-US")
			newsReq.Header.Set("User-Agent", "Mozilla/5.0")

			newsResp, _ := client.Do(newsReq)
			newsBytes, _ := ioutil.ReadAll(newsResp.Body)
			newsResp.Body.Close()

			var news News
			xml.Unmarshal(newsBytes, &news)

			// Store the article
			newsItems[location] = news

		}
	}

	// Parse the template file
	t, _ := template.ParseFiles("news-item.html")

	// Use Execute err to uncover problems in the template page
	fmt.Println(t.Execute(w, newsItems))
}

// The / page handler
func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Yo! I made a server</h1>")
}

func main() {

	// Start web server
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}
