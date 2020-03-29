package main

import ("fmt" 
		"net/http"
	)

func index_handler(w http.ResponseWriter, r *http.Request){
	//unused template
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hey there!")
}

func main(){
	//creating a handlers that would figure out what kind of function
	//corresponds to that path
	//index page
	http.HandleFunc("/", index_handler)
	//about page
	http.HandleFunc("/about/", about_handler)
	//have server, nil = none or null
	http.ListenAndServe(":8000", nil)

	
	 
}