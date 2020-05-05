package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	// to make sure that we firstly finish add all values into the chennel.
	defer wg.Done()
	// send some value time 5 to the channel
	c <- someValue * 5
}

func main_for_go_channel() {
	// make a channel with type int
	// assign to some variable because if you want to be able to parse make to other types.
	fooVal := make(chan int)

	// create channel and assign some value.
	go foo(fooVal, 5)
	go foo(fooVal, 3)

	//receive the value to the channel
	// we call to receive values twice, because we know that we have 2 channels.
	// what if we don't know exactly number of channels we have.
	// we had better to use something instead of this sloppy way.
	v1 := <- fooVal
	v2 := <- fooVal
	fmt.Print(v1,v2) //15 25
	// or
	// v1 ,v2 := <-fooVal, <-fooVal //15 25

}

func main_for_iterating_channel(){
	// we have to buffer for 10 items then add 10.
	fooVal := make(chan int, 10)
	// create i channels and assign value i.
	for i := 0; i < 10; i++ {
		// method foo with i = 0 = queue 1, then 2, ...
		wg.Add(1)
		go foo(fooVal, i)
	}
	// to prevent close the channel while all/some values don't go inside. 
	wg.Wait()
	close(fooVal) // fix the error, but the result is nothing.

	// print items in the channel
	for item := range fooVal {
		fmt.Println(item) 
	}
}

func main(){
	// main_for_go_channel()
	main_for_iterating_channel()
}