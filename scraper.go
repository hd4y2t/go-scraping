package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deksripsi"`
	Img       string `json:"img"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("kesbangpolpalembang.com"),
	)

	c.OnHTML("section[class=course] div.info-card h6", func(e *colly.HTMLElement) {
		link := e.Text
		fmt.Println(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://kesbangpolpalembang.com/data-artikel/")
}
