package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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

3.

*/

// <sitemap>
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// <loc>
type Location struct {
	Loc string `xml:"loc"`
}

// Provide a neater print of the Location
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	client := &http.Client{}

	// GET from upstream feed
	// Necessary to have headers set to avoid tar pit from upstream
	req, _ := http.NewRequest("GET", "https://www.washingtonpost.com/news-sitemaps/index.xml", nil)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	// Parse the body as XML using Location type
	var sitemapIndex SitemapIndex
	xml.Unmarshal(bytes, &sitemapIndex)

	fmt.Println(sitemapIndex.Locations)

	// Iterate over Locations using range (index and element)
	for index, Location := range sitemapIndex.Locations {
		fmt.Printf("Index: %v %s", index, Location)
	}

}
