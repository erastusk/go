package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

var ScrapeUrl = "https://quotes.toscrape.com/"

type QuotesOfTheDay struct {
	Mood   string
	Quotes []string
}

func getTags() []string {
	tags := []string{}
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnHTML(".tag-item a", func(resp *colly.HTMLElement) {
		tags = append(tags, resp.Text)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.Visit(ScrapeUrl)
	return tags

}

func ReadQuotesFromTag(url string) []string {
	quotes := []string{}
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnHTML(".text", func(resp *colly.HTMLElement) {
		quotes = append(quotes, resp.Text)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.Visit(url)
	return quotes

}

func getQuotes(s string) []string {
	url := "https://quotes.toscrape.com/tag/" + s + "/"
	quotes := ReadQuotesFromTag(url)

	return quotes
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple web app")
}
func quotes(w http.ResponseWriter, r *http.Request) {
	retTags := getTags()
	rand.Seed(time.Now().Unix())
	TagOfTheDay := retTags[rand.Intn(len(retTags))]
	bodyHtml := getQuotes(TagOfTheDay)
	p := QuotesOfTheDay{Mood: TagOfTheDay, Quotes: bodyHtml}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)

}

func main() {

	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/quotes", quotes)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
