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

### Sample Outputs:

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

 ```

 You can pass the site you want to create the sitemap
 ```
> go run main.go -depth=3 -outFile="go-map.xml" -url="https://example.com/"
2021/11/14 18:41:25 Sitemap for 'https://example.com/' is written to 'go-map.xml' successfully 
>
 ```

 Please have a look at the `main.go` for a working example.

 **Example:**
Generated sitemap.xml from above library
```
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://www.pybuzz.com/running-python-programs-in-linux/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/comments-in-python-programming-language/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/rules-to-name-identifiers-or-variables-in-the-python-programming-language/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/intro/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/basics/page/2/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/compiler-vs-interpreter/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/running-python-programs-in-linux/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/installing-python-3-on-windows-10-7/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/author/venkey/page/2/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#1-what-is-python-interactive-mode</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/installing-python-3-on-windows-10-7/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-indentation/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/tag/basics/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/tag/building-blocks/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/basics/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/keywords-and-identifiers-in-python-programming-language/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#3-python-script-mode</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#0-introduction</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/writing-your-first-python-program/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/installing-python-3-on-ubuntu-linux/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-programming-language-and-features/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#4-example-calculator-program-in-python</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/tag/python/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/rules-to-name-identifiers-or-variables-in-the-python-programming-language/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/statements-in-python-programming-language/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/tag/intro/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/compiler-vs-interpreter/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-programming-language-and-features/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/page/2/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-indentation/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/writing-your-first-python-program/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#2-example-interactive-python-statements</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/page/2/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/statements-in-python-programming-language/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/introduction-to-programming-language/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/author/venkey/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#6-conclusion</loc>
  </url>
  <url>
    <loc>http://www.pybuzz.com/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/keywords-and-identifiers-in-python-programming-language/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/comments-in-python-programming-language/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/installing-python-3-on-ubuntu-linux/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/python-interactive-mode-and-script-mode/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/introduction-to-programming-language/#respond</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/</loc>
  </url>
  <url>
    <loc>https://www.pybuzz.com/category/python/building-blocks/</loc>
  </url>
</urlset>
```