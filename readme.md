This Go Module will generate the sitemap for any given URL.

We will travese the all pages of the website and creates the sitemap

We will use the
- 'net/http' module to get the webpage
- 'https://github.com/mvenkatesh431/LinkParser' to parse the links on a WebPage

We will filter out the links(Ignoring the fragments and links to other sites, etc) and use the BFS to traverse through the pages to get all links.

Note: This is my version of one of the excersices of *gophercise*


Then we will generate the XML sitemap using above links. 
We will generate the XML format as per the standard sitemap protocol
```
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```

### Sample Output:

**Usage:**

```
> go run main.go -h
Usage of C:\Users\mvenk\AppData\Local\Temp\go-build1601473586\b001\exe\main.exe:
  -depth int
        The Max depth of pages you can Traverse (default 5)
  -outFile string
        Sitemap will be saved this file (default "sitemap.xml")
  -url string
        Website URL to create the Sitemap (default "http://www.pybuzz.com/")
>


> go run main.go -depth=3 -outFile="go-map.xml" -url="https://example.com/"
2021/11/14 18:41:25 Sitemap for 'https://example.com/' is written to 'go-map.xml' successfully 
>

 ```