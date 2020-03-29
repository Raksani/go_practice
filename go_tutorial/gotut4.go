package main

import ("encoding/xml"
	"fmt"
	"net/http"
	"io/ioutil"
)
/*
	for web(more)

func index_handler1(w http.ResponseWriter, r *http.Request) {
	//do thi instead of below because when debug it, it will appear as multi line.

	fmt.Fprintf(w,` <h1>Hey there</h1>
	<p>Go is fast </p>
	<p>You are <strong>bold</strong></p>
	`)
}
func main(){
 	http.HandleFunc("/", index_handler1)
 	http.ListenAndServe(":8000",nil)	
 }
*/ 

/*
	 From the xml file, <sitemapindex> is the parent.
	 Then, we create a slice, called Locations and add the `xml:"sitemap"`
	 to understand where it is looking when we go to unpack this with encoding/xml package
*/
type SitemapIndex struct {
	//Locations as a array of Location type
	//L in Location must be capitalized
	/*
		[5 5]type = array
		[]type = slice 
		different? array has fixed size (5x5)
		slice is just, for example, Locations is just slice of Location
	*/
	Locations []Location `xml:"sitemap"`
}


/*
	Then, building a slice, called Locations which has type Location. 
	พอได้แกะsitemapออกมาใส่ไว้ในslice ของ location เราก็ทำstruct location ต่อเพื่อแกะ <loc>
	ในlocation type ก็จะประกอบด้วย Loc variables ที่เป็นstring
	และก็มี xml:"loc"` บอก locationของtagที่เราใช้สำหรับLoc

*/

type Location struct {
	Loc string `xml:"loc"`
}

/*
	After we get the result, there are {url} in slice instead of url
	because these URLs is Location type which has Loc elements as string
	So, we want strings not Location type by  building a string method.
	*Note that we have no need to call this method in main, I personally
	understand that because it is as same as overwrite the existing String() method.
	So, any named type (l Location) will be custom as string (not sure yet)
*/
func (l Location) String () string {
	//Sprintf will return string 
	return fmt.Sprintf(l.Loc)
  }

// goal: Creating a news aggregator web app 
// task: read an information from the internet (pull data)
func main(){
	// make a request and then the request will return response in bytes
	// and error. An error might be empty, for now we define it as underscore
	// we can define a variable that you don't intend to use as _ (underscore)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	//unpack it
	bytes, _ := ioutil.ReadAll(resp.Body)
	// parse to string
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s) // memory address s

	// fmt.Println(s.Locations)
	// After we know how to loop, we can print all the URLs as multilines of strings
	fmt.Printf("Here %s some %s", "are","variables")
	// we use _, Location because in map in will contain both key and value.
	// then _ is the index of the url (key) which we don't want to print and know. 
	// So, we use _ that we already declared (just to iterate like i or something)
	// what we really want is the next one which is the "value" or the real url data which uses Location as iterator
	// *Note that _, Location = 1, http://...
	// *and Location is just a variable which is not associated w/ Location struct.
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}

