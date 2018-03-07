package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

var hasErrors = false

func init() {
	log.SetFlags(0)
	log.SetPrefix("lint: ")
}

func main() {
	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatalln(err)
	}

	html := blackfriday.MarkdownCommon(readme)
	buf := bytes.NewBuffer(html)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		log.Fatalln(err)
	}

	var currentHeading string

	doc.Find("h2, h3, ul").Each(func(i int, s *goquery.Selection) {
		switch {
		case s.Is("h2, h3"):
			switch s.Text() {
			case "Awesome Hetzner Cloud":
				// Do not check contents list
				currentHeading = ""
				return
			default:
				currentHeading = s.Text()
			}
		case s.Is("ul"):
			if currentHeading != "" {
				checkList(s, currentHeading)
			}
		}
	})

	if hasErrors {
		os.Exit(1)
	}
}

func checkList(ul *goquery.Selection, heading string) {
	var lastName string
	ul.Find("li > a:first").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		if i == 0 {
			lastName = name
			return
		}
		if strings.ToLower(name) < strings.ToLower(lastName) {
			fmt.Printf("item %q in list %q is not sorted alphabetically (must be before %q)\n", name, heading, lastName)
			hasErrors = true
		}
		lastName = name
	})

	ul.Find("li").Each(func(i int, s *goquery.Selection) {
		if !strings.Contains(s.Text(), "—") {
			name := s.Find("a").Text()
			fmt.Printf("item %q does not use correct dash (—)\n", name)
			hasErrors = true
		}
	})
}
