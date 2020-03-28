package main
import ("fmt" 
		"math/rand"
		"math")

func add(x float64,y float64) float64{
	return x+y
}

func multiple(a,b string) (string,string){
	return a,b
}

func main() {
	fmt.Println("The square root of 4 is ",math.Sqrt(4))
	
	//call the method random once, it will truly honest with you until you change the seed (or call this method as a second time)
	//So, when you use go run this file, the random will always be the same every time.
	fmt.Println("A number between 0-100 is ", rand.Intn(100)) //81
	fmt.Println("A number between 0-100 is ", rand.Intn(100)) //87
	
	var num1,num2 float64 = 5.6,9.5
	//make them be public static, and go will decide which types they are.
	//but make sure that your func can use this type.
	//for example, func add float32. But num3,4 by go basically n as float64 as default.
	//So, they cannot be used in this function.
	num3,num4 := 5.0, 10.0
	fmt.Println(add(num1,num2))
	fmt.Println(add(num3,num4))

	w1,w2 := "Hey", "there"
	fmt.Println(multiple(w1,w2))

	//Pointer
	x := 15
	a := &x //memory address of x
	fmt.Println(a) //print 0xc0000b4030 (address of x)
	fmt.Println(*a) // print value in the address a which is 15
	*a = 5 //change the value in address "a" to be 5
	fmt.Println(x) // print the current value in the address 'a' 
					//which is now "5"
	*a = *a**a //value in address "a" (5) equals to
				//value in address "a" (5) * value in address "a" (5)
				//the result is 25!
	fmt.Println(*a)
	fmt.Println(x) //as x is the value in the address "a" as well 
	
}
