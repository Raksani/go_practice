package main

import ("fmt")

func foo() {
	// //it prints "done" when the func foo is finished
	// defer fmt.Println("done") //third (wait for third)
	// defer fmt.Println("Are we done?") //second (wait for one)
	// fmt.Println("doing some stuff, who knows what?") //first
	
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	} //it will print 43210 because it wait until it finished (first in, last out)

	
}

func main() {
	foo()
}