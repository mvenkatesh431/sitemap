package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mvenkatesh431/LinkParser/link"
)

func main() {

	urlVar := flag.String("url", "http://www.pybuzz.com/", "Website URL to create the Sitemap")
	flag.Parse()

	siteLinks, err := get(*urlVar)
	if err != nil {
		log.Fatalf("Failed to get the links %v\n", err)
	}

	fmt.Println(siteLinks)
}

/*
	The 'get' function will get the 'urlVar' webpage using the 'http.get'
	Then we will parse the all the links using the 'LinkParser'
*/
func get(urlVar string) ([]string, error) {
	resp, err := http.Get(urlVar)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// This will help us to get the base url, Even though user provides the some page link(example.com/about-us)
	baseUrl := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host

	return (getHrefs(resp.Body, baseUrl))
}

/*
	The 'getHrefs' will take the 'io.Reader'(resp.Body) and baseUrl
	Parses the links and returns the slice of strings(links)
*/
func getHrefs(reader io.Reader, baseUrl string) ([]string, error) {

	// Parse the links from the webpage using 'https://github.com/mvenkatesh431/LinkParser'
	links, err := link.GetLinks(reader)
	if err != nil {
		return nil, err
	}

	var siteLinks []string

	// Lets remove the fragments and create absolute links.
	for _, link := range links {

		switch {
		case strings.HasPrefix(link.Href, "/"):
			// This case covers all links which are starting with "/"
			// We need to add the 'baseUrl' to this 'Href'
			siteLinks = append(siteLinks, baseUrl+link.Href)

		case strings.HasPrefix(link.Href, "http"):
			// This will filter all valid links(http and https)
			if strings.HasPrefix(link.Href, baseUrl) {
				// SiteMap contains baseUrl links only, So add only 'baseUrl' links to sitemap
				siteLinks = append(siteLinks, link.Href)
			}
		}
	}
	return siteLinks, nil
}
