package main

import (
	"fmt"
	// "unsafe"
)

const secondsInHour = 3600

func main() {

	var age int = 21
	var name string = "Iman"

	fmt.Println("Hello Worlds")
	fmt.Println("your age is: ", age)
	fmt.Println("next year your age is: ", age+1)
	fmt.Println("your name is : ", name)

	country := "Iran"
	fmt.Println("your country is : ", country)

	var isStudent bool = true
	fmt.Println("is sudent or not: ", isStudent)

	var pi float64 = 3.14
	fmt.Println("the value of pi is : ", pi)

	fmt.Println("the value of pi * 2 is : ", pi*2)

}
