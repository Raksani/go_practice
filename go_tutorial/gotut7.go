package main

import ("fmt"
		"strings"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
	)

// we create struct to pass multiple values to our template. Since the index_handler 
// did not use any template just a line of HTML 
type NewsAggPage struct{
	Title string
	// News string
	News map[string]NewsMap
}

/*
	from the news aggregator application (gotut5.go)
*/

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	//Titles: key of the map
	//Keywords and Locations: Value (Location)
	Keyword string
	Location string
}
/*
	From the struct, we want to put value on it
	p is our new type, NewsAggPage with value of Title and News
	t, _ is our template

*/
func newsAggHandler(w http.ResponseWriter, r *http.Request){

	fmt.Print("newsAggHandler processing")
	/*
		from the news aggregator application (gotut5.go)
	*/
	var s SitemapIndex
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s) // memory address s
	
	// fmt.Printf("Here %s some %s", "are","variables")
	
	// for _, Location := range s.Locations {
	// 	fmt.Printf("\n%s", Location)
	// }

	var n News

	//create a map
	
	
	for _, Location := range s.Locations {
		//After we ran our program, we got an error about runtime error
		//the problem is related to extra spaces in xml locations 
		//which has shown when we looping and printing the URLs out there are spaces as shown
		//So, we have to fix those by get rid of spaces
		Location = strings.TrimSpace(Location)
		//So, we can get EACH of url in string(Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)

		for i, _ := range n.Keywords {
			newsMap[n.Titles[i]] = NewsMap{n.Keywords[i], n.Locations[i]} 
		}
	}
	// p is to add "Amazing News Aggregator" as Title and content inside that news as News
	// p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}

	// But we want our news content instead of "some news". So, we should iterate our NewsMapAgg type
	// which will basically be passed to our document in the newsAggPage type, under News name
	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	//in the HTML file, we use News variable as {{ .News}}

	
	// t, _ := template.ParseFiles("basictemplating.html")
	/*
		basictemplating.html
			<h1>{{.Title}}</h1>
			<p>{{.News}}</p>
	*/

	//change template
	// t, _ := template.ParseFiles("newsaggtemplate.html")

	//change template
	t, _ := template.ParseFiles("aggregatorfinish.html")

	t.Execute(w, p)

}

func main(){
	//creating a handlers that would figure out what kind of function
	//corresponds to that path
	
	http.HandleFunc("/agg/", newsAggHandler)
	
	//have server, nil = none or null
	http.ListenAndServe(":8000", nil)
}