package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(link, ".pdf") && strings.Contains(link, "siamak.page") {
			c.Visit(link)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r* colly.Response) {
		link := r.Request.URL.String()
		fmt.Printf("Link: %s\n", link)
		if strings.Contains(link, ".pdf") {
			fmt.Println("PDF found! ", r.Request.URL)
			s := strings.Split(link, "/")
			name := s[len(s)-1]
			r.Save(name)
		}
	})

	c.Visit("https://www.siamak.page/courses/COMP551F20/index.html")
}