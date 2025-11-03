package main

import (
	"fmt"
)

func main() {
	age := 2

	if age >= 18 {
		fmt.Println("your are bigger than 18 years old")
	} else {
		fmt.Println("your are too young")
	}



	// for loops
	for i:=1; i<=10; i++ {
		fmt.Println("the i is :" , i)
	} 


	// while loop
	count := 20
	for count >= 10 {
		count --
		if count == 18 {
			continue
		}
		if count == 15 {
			break
		}

		fmt.Println("the count is : ", count)	
		fmt.Println("---------------------------")
	}

	// switch case
	var day string
	fmt.Println("Enter day : ")
	fmt.Scan(&day)
	switch day {
	case "monday":
		fmt.Println("today is monday")

	case "sunday":
		fmt.Println("today is sunday")
	
	case "friday":
		fmt.Println("today is friday")	
	
	default:
		fmt.Println("ok")
	
	}
	fmt.Println("-------------------------------------")
	fmt.Println()

	// functions
	Great()
	Great()

	yearlySalary(12000)

	fmt.Println(add(3,2.3))

}

func Great(){
	fmt.Println("hello and welcome")
}


func yearlySalary(salary int64) {
	fmt.Println("your yearly salary is: ", salary * 12)
}

func add(a float32, b float32) float32 {
    return a + b
}