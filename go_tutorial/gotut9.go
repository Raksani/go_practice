package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup //wg = wait group

func say(s string) {
	defer wg.Done() //if below func has error/panics out, this defer will do as recovery, 
	// means that it will be called for sure.
	// but why we want the program to be recovered? because when something hit an error panic out
	// generally the program would end, but we want it keep going.
	for i :=0; i<=2; i++ {
		// fmt.Println(s)
		// time.Sleep(time.Microsecond*100)
		time.Sleep(time.Microsecond*100)
		fmt.Println(s)
	}
	// wg.Done() //to notify the wait group, it done.
}

func main(){
	//before the start
	wg.Add(1) //there is 1 in a wait group
	go say("Hey")
	wg.Add(1) //there is 1 in a wait group
	go say("There")
	wg.Wait() //wait until go1 and 2 are finished
	// But anyway, if wg.Wait() until wg.Done and something wrong w/ the loop, wg.Done()
	// have not been called for a long time and wg.Wait() wait for...ever
	// since the problem we might come accross, we have solution for that -> defer statement

}
