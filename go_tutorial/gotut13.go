/*
	Modify by using sync (channel, wait) --> run faster
*/

package main

import ("fmt"
		"strings"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
		// 1: import
		"sync"
	)

// 2: define to create a wait group
var wg sync.WaitGroup

type NewsAggPage struct{
	Title string
	News map[string]NewsMap
}


type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

// 3: create new routine to push all the data in side the channel
// but what is our channel going to contain? anything in type 'news'
// So, we can either pass n to the channel or create local News per channel
// Then, we should remove var n News from newsAdggHandler and place it here. 
func newsRoutine(c chan News, Location string) {
	// 12: defer the done wait group
	defer wg.Done()

	// remove var n News from newsAdggHandler
	var n News
	//method from for-loop in newsAggHandler
	Location = strings.TrimSpace(Location)
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	//send the value 'n' to the channel.
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){

	fmt.Print("newsAggHandler processing")
	var s SitemapIndex
	newsMap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s) 

	// 3: being moved to newsRoutine()
	// var n News

	// 4: create a channel to pass data from running goroutines through the
	//    newsRoutine call. 
	queue := make(chan News, 30)
	for _, Location := range s.Locations {
		// 10: to make a wait for each newsRoutine below.
		wg.Add(1)
		// 5: make newsRoutine a goroutine.
		go newsRoutine(queue, Location)

		// 6: take these below out. Because that's going to send it over to
		//    the channel. So, move it out of this loop.
		// for i, _ := range n.Keywords {
		// 	newsMap[n.Titles[i]] = NewsMap{n.Keywords[i], n.Locations[i]} 
		// }
	}

	// 11: After we run all of the goroutines then close the channel and 
	//     it will continue the iteration range as below.
	wg.Wait()
	close(queue)


	// 8: Continue from step7. So, we need to iterate over the channel first and
	//    then we've got this 'News' type basically. 
	// elem = 'that n we mentioned' 
	// 9: replace 'n' with 'elem'
	for elem := range queue { 	
		// 7: we're gonna want to do this iteration over n.Keywords... and stuffs.
		//    but we don't just have a single 'n'as before, we've cluster of 'n'. 
		//	  and this cluster of 'n' is coming over each channel. So, --> step 8!
		// for i, _ := range n.Keywords {
		for i, _ := range elem.Keywords {
				// newsMap[n.Titles[i]] = NewsMap{n.Keywords[i], n.Locations[i]}
				newsMap[elem.Titles[i]] = NewsMap{elem.Keywords[i], elem.Locations[i]} 
			}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("aggregatorfinish.html")

	t.Execute(w, p)

}

func main(){
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}