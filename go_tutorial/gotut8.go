package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i :=0; i<=3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond*100)
	}
}

func main(){
	go say("Hey")
	go say("There")
	
	time.Sleep(time.Second*2) //to wait for Hey finish, then there.
}