package sitemap

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mvenkatesh431/LinkParser/link"
)

type emptyStruct struct{}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// The CreateSitemap function will take the "website URL", "Depth", "output sitemap file" and creates the sitemap.
func CreateSitemap(urlVar string, maxDepth int, outFile string) error {
	// The 'parseLinks' will do the breadth first parseLinks and goes to all the pages
	siteLinks := parseLinks(urlVar, maxDepth)
	return xmlSitemap(siteLinks, outFile)
}

/*
	The 'xmlSitemap' will takes the siteLinks(all Links) as a slice of strings
	and create a sitemap in the sitemap standard format and saves the sitemap into "outFile"
*/
func xmlSitemap(siteLinks []string, outFile string) error {
	encXml := urlset{
		Xmlns: xmlns,
	}

	for _, link := range siteLinks {
		encXml.Urls = append(encXml.Urls, loc{link})
	}

	// Open file for writing and truncate if there is data
	f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(xml.Header)
	encoder := xml.NewEncoder(f)

	// will have two spaces as indentation
	encoder.Indent("", "  ")
	if err := encoder.Encode(encXml); err != nil {
		return err
	}
	return nil
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

/*
	The 'parseLinks' will do the breadth first parseLinks and goes to all the pages
	'urlVar' is the url given by user
	'parseLinks' returns the slices of string - which contains all the links in the website
*/
func parseLinks(urlVar string, maxDepth int) []string {

	var result []string

	// We don't need to go through the pages we already seen. So maintain a map and check before parsing
	// Our 'seen' map contains the string and empty struct. The key will be the 'links'
	// The empty struct 'struct{}' won't assign any memory, But you can also use 'bool' type.
	// lets define emptyStruct type and use it.
	seen := make(map[string]emptyStruct)

	// We need two maps, One will contain the links we are going through now.
	// Other will contian the links we found in this iteration
	var present map[string]emptyStruct
	future := map[string]emptyStruct{
		urlVar: emptyStruct{}, // add the 'urlVar'(our base link) to 'future' map.
	}

	// Lets iterate upto 'MaxDepth'
	for i := 0; i <= maxDepth; i++ {

		// copy 'future' to 'present' and make 'future' empty
		present, future = future, make(map[string]emptyStruct)

		// if the 'present' map is empty, Means we reached end. So break it.
		if len(present) == 0 {
			break
		}

		for page, _ := range present {
			// if the 'page' exists in the 'seen', Then we already went to through page, So 'continue'
			if _, ok := seen[page]; ok {
				continue
			}
			seen[page] = emptyStruct{}

			links, err := get(page)
			if err != nil {
				// If we failed to get a 'page', Then continue.
				continue
			}

			// Now call the 'get' to parse all links
			for _, link := range links {
				// Add the 'link' to 'future' map, So that we can visit them later
				if _, ok := seen[link]; !ok {
					// only add the links which are not present in 'seen' map
					future[link] = emptyStruct{}
				}
			}
		}
	}

	// We need to return the slice of strings, so populate the 'result' and return
	for url, _ := range seen {
		result = append(result, url)
	}

	return result
}
