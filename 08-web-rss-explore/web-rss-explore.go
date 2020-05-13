package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*

1. GET from the upstream server

<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
...
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/world.xml
</loc>
</sitemap>
</sitemapindex>

2. Parse the XML to extract the additional links

3. GET an upstream link and parse its XML

<urlset>
<url>
<loc>
https://www.washingtonpost.com/politics/2020/05/12/fauci-testimony-senate-coronavirus/
</loc>
<lastmod>2020-05-12T16:41:49.599Z</lastmod>
<news:news>
<news:publication>
<news:name>Washington Post</news:name>
<news:language>en</news:language>
</news:publication>
<news:publication_date>2020-05-12T16:41:49.599Z</news:publication_date>
<news:title>
Fauci warns Senate that reopening U.S. too quickly could lead to avoidable ‘suffering and death’
</news:title>
<news:keywords>
fauci, fauci testimony senate, fauci testify senate
</news:keywords>
</news:news>
<changefreq>hourly</changefreq>
</url>
...
</urlset>
*/

// Refactored into a single "master" struct <sitemap>...<loc>
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// The "slave" sitemap has a more detailed XML schema
//
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {

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

	// Iterate over sitemap locations using range (index and element)
	// and put the results in a map

	// Create a map to store the news items
	newsItems := make(map[string]News)

	for _, location := range sitemapIndex.Locations {

		// Location may have whitespace
		location = strings.TrimSpace(location)

		fmt.Printf("\nConsider: %s\n", location)

		// Only interested in politics and opinions to keep traffic down
		if strings.Contains(location, "politics.xml") ||
			strings.Contains(location, "opinions.xml") {

			// Print the URL
			fmt.Printf("GET: %s", location)

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

			newsItems[location] = news

		}
	}

	for key, item := range newsItems {
		fmt.Printf("\nKey: %s\n", key)
		fmt.Printf("\nTitle: %s", item.Titles[0])
		fmt.Printf("\nKeywords: %s", item.Keywords[0])
		fmt.Printf("\nLocations: %s", item.Locations[0])
	}
}
