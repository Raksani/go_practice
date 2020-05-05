// Panic: to make the program disability and start to panic.
// This will STOP RUNNING FUNCTION and RUN all of the DEFERed functions from
// the panicking function. 
// Recover: program had stopped running by panic which is not good.
// So, we recovered it from a panicking gooroutine. 
// you have to put RECOVER into one of these DEFERed functions.

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup(){
	// handle panic by checking for a recover.
	// Also, we would like to save recover() as r. So, we can use r later.
	// as below we  print r result ( r = 'Oh dear... a 2')
	// we pass this sentence below to PANIC
	if r := recover(); r != nil{
		fmt.Println("Recovered in cleanup: ", r) // Recovered in cleanup Oh dear... a 2
	}

	wg.Done() //move to here, same result. Because defers go in reverse order that they come in.
	// So, just order it like this is better.
}

func say(s string) {
	// defer wg.Done() // is the function we're deferring, but this is not our function.
	// So, we have to create the new one above in the name cleanup(). 

	defer cleanup();

	for i := 0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
		// if i == 2, it will definitely generate error
		if i == 2 {
			panic("Oh dear... a 2") // the result is panic = everything falling apart --> don't want it to be like this.
		}
	}	
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Wait()
}
