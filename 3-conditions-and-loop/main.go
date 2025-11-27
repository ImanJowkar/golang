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



	// for loops
	for i:=1; i<=10; i++ {
		fmt.Println("the i is :" , i)
	} 

	// for loops
	i := 0
	for i < 10 {
		fmt.Println("the i is :" , i)
		i++
	} 

	// for on list
	lst := []int{4,5,6,7,3,2,1,0}
	for idx, val := range lst {
		println("idx is : ",idx, " and value is : ",val)
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



}

