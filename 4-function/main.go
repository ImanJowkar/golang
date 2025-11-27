package main

import (
	"fmt"
)

func main() {

	// functions
	Great()
	Great()

	yearlySalary(12000)

	fmt.Println(add(3,2.3))

	var res float64
	res = calc(4.3, 34)
	fmt.Printf("The value of res is : %f", res)
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


func calc(a float64, b float64) (result float64) {
	var operation string
	fmt.Println("What do you want : sum, divide, mines , multiply? ")
	fmt.Scan(&operation)

	switch operation {
	case "sum":
		result = a + b

	case "divide":
		result = a / b
	
	case "mines":
		result = a - b

	case "multiply":
		result = a * b
		
	default:
		fmt.Println("not defined")
	
	}
	return result
 
}

