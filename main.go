package main

import (
	"flag"
	"log"

	"github.com/mvenkatesh431/sitemap"
)

func main() {

	urlVar := flag.String("url", "http://www.pybuzz.com/", "Website URL to create the Sitemap")
	maxDepth := flag.Int("depth", 5, "The Max depth of pages you can Traverse")
	outFile := flag.String("outFile", "sitemap.xml", "Sitemap will be saved this file")
	flag.Parse()

	// The CreateSitemap function will take the "website URL", "Depth", "output sitemap file" and creates the sitemap.
	err := sitemap.CreateSitemap(*urlVar, *maxDepth, *outFile)
	if err != nil {
		log.Fatalf("Failed to create the XML Sitemap %v\n", err)
	}
	log.Printf("Sitemap for '%s' is written to '%s' successfully \n", *urlVar, *outFile)
}
