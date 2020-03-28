package main

import "fmt"

func main() {
	//key as string 
	//value as float32
	grades := make(map[string]float32)

	grades["Beau"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	BeausGrade := grades["Beau"]
	fmt.Println(BeausGrade) //42

	//we can delete from where, who
	delete(grades, "Beau") //Beau no longer in the map

	for k, v := range grades{
		fmt.Println(k ,":", v) //Jess : 92
	}

	//we use these knowledge applying to gotut5.go

}


