This Go Module will generate the sitemap for any given URL.

We will travese the all pages of the website and creates the sitemap

We will use the
- 'http/net' module to get the webpage
- 'https://github.com/mvenkatesh431/LinkParser' to parse the links on a WebPage

We will filter out the links(Ignoring the fragments and links to other sites, etc) and create the sitemap.

Note: This is my version of one of the excersices of *gophercise*

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