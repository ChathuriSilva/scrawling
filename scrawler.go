package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Enter Item Type:")
	var item string
	fmt.Scanln(&item)

	fmt.Println("Enter District:")
	var dis string
	fmt.Scanln(&dis)

	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk", "https://ikman.lk/en/ads"),
	)
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// 	c.Visit(e.Request.AbsoluteURL(link))
	// })

	c.OnHTML(".list--3NxGO", func(e *colly.HTMLElement) {
		title := e.Attr("class")
		fmt.Printf("Title Found: %q -> %s\n", e.Text, title)

		// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// 	link := e.Attr("href")
		// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// c.OnHTML(".word-break--2nyVq", func(e *colly.HTMLElement) {
		// 	information := e.Attr("class")
		// 	fmt.Printf("Data Found: %q -> %s\n", e.Text, information)
		// 	c.Visit(e.Request.AbsoluteURL(information))
		// })

		// 	c.Visit(e.Request.AbsoluteURL(link))
		// })
		c.Visit(e.Request.AbsoluteURL(title))
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://ikman.lk/en/ads/" + dis + "/" + item)
}
