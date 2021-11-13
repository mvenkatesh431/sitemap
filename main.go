package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mvenkatesh431/LinkParser/link"
)

func main() {

	urlVar := flag.String("url", "http://www.pybuzz.com/", "Website URL to create the Sitemap")
	flag.Parse()

	fmt.Printf("Site URL : %v\n", *urlVar)

	resp, err := http.Get(*urlVar)
	if err != nil {
		log.Fatalf("Failed to retrive the webpage %v\n", err)
	}
	defer resp.Body.Close()

	links, err := link.GetLinks(resp.Body)
	if err != nil {
		log.Fatalf("Failed to get the links %v\n", err)
	}

	// This will help us to get the base url, Even though user provides the some page link(example.com/about-us)
	baseUrl := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host

	var siteLinks []string

	// Lets remove the fragments and create absolute links.
	for _, link := range links {

		switch {
		case strings.HasPrefix(link.Href, "/"):
			// This case covers all links which are starting with "/"
			// We need to add the 'baseUrl' to this 'Href'
			siteLinks = append(siteLinks, baseUrl+link.Href)

		case strings.HasPrefix(link.Href, "http"):
			// This will filter all valid links.
			if strings.HasPrefix(link.Href, baseUrl) {
				siteLinks = append(siteLinks, link.Href)
			}
		}
	}

	fmt.Println(siteLinks)
}
