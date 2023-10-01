package main

import (
	"encoding/json"
	"fmt"
	"os"

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

	var items []item

	c.OnHTML("section[class=course]", func(e *colly.HTMLElement) {
		item := item{
			Nama:      e.ChildText("h6"),
			Deskripsi: e.ChildText("h2"),
			Img:       e.ChildAttr("img", "src"),
		}
		items = append(items, item)
	})

	c.OnHTML("[arial-label=Next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://kesbangpolpalembang.com/data-artikel")
	content, err := json.MarshalIndent(items, "", " ")

	if err != nil {
		fmt.Println("Error", err)
	}

	os.WriteFile("data.json", content, 0644)

}
