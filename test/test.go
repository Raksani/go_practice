package main

import (
	"strings"
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
)


type NewsMap struct {
	Keyword string
	Location string
}

type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("newsAgg")
	var s Sitemapindex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		fmt.Printf("for range Locations")
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			fmt.Printf("for range Keywords")
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
			fmt.Print(news_map[n.Titles[idx]])
		}
	}
	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
    t, _ := template.ParseFiles("aggregatorfinish.html")
	t.Execute(w, p)
	fmt.Print("executed")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil) 
}