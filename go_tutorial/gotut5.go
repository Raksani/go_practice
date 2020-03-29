package main

import ("strings"
	"fmt"
	"encoding/xml"
	"net/http"
	"io/ioutil"
)

type SitemapIndex2 struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords[]string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	//Titles: key of the map
	//Keywords and Locations: Value (Location)
	Keyword string
	Location string
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex2
	xml.Unmarshal(bytes, &s) // memory address s
	
	// fmt.Printf("Here %s some %s", "are","variables")
	
	// for _, Location := range s.Locations {
	// 	fmt.Printf("\n%s", Location)
	// }

	var n News

	//create a map
	news_map := make(map[string]NewsMap)
	
	for _, Location := range s.Locations {
		//After we ran our program, we got an error about runtime error
		//the problem is related to extra spaces in xml locations 
		//which has shown when we looping and printing the URLs out there are spaces as shown
		//So, we have to fix those by get rid of spaces
		Location = strings.TrimSpace(Location)
		//So, we can get EACH of url in string(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for i, _ := range n.Keywords {
			news_map[n.Titles[i]] = NewsMap{n.Keywords[i], n.Locations[i]} 
		}
	}
	for j, data := range news_map {
		fmt.Println("\n\n\n\n",j)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}	
}