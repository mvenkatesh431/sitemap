package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mvenkatesh431/LinkParser/link"
)

func main() {

	urlVar := flag.String("url", "https://pybuzz.com", "Website URL to create the Sitemap")
	flag.Parse()

	fmt.Printf("Site URL : %v\n", *urlVar)

	resp, err := http.Get(*urlVar)
	if err != nil {
		log.Fatalf("Failed to retrive the webpage %v\n", err)
		// handle error
	}
	defer resp.Body.Close()

	links, err := link.GetLinks(resp.Body)
	if err != nil {
		log.Fatalf("Failed to get the links %v\n", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}

}
