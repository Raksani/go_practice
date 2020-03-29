package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // 0 to 65535
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

//method to calculate km per hr
// c is associate w/ car struct
// (c car) = func receive c car type
//func kmh() returns the float64 
//like parse the car type to do some operations

func (c car) kmh() float64 {
	//parse float64 for gas_pedal because top speed/u is float64.
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}
//function to change variable's value in the struct
func (c *car) new_top_speed(newspeed float64){
	c.top_speed_kmh = newspeed
}
func main(){
	//readable way 
	a_car := car{gas_pedal: 22341, 
		brake_pedal: 0, 
		steering_wheel: 12561, 
		top_speed_kmh: 225.0}
	
	// b_car := car{22341,0,12562,225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(500)
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())

}
