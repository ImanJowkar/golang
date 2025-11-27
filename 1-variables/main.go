package main

import (
	"fmt"
	// "unsafe"
)

const secondsInHour = 3600

//struct
type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

func main() {
	//var name string = value	
	var city string
	var province string
	city = "shiraz"
	province = "fars"

	fmt.Println("city and province", city, province)
	
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


	// data type can found automatically
	e := 2.71
	fmt.Println("the e is: ", e)

	name1 := "saman"
	fmt.Println("the name1 is: ", name1)

	var i1, i2, i3 = 1,2,3
	fmt.Println(i1,i2,i3)

	// const data types
	const (
		youtubeName = "Iman Jowkar"
		cityBorn = "Shiraz"
	)

	var (
		employeeName string = "Jack"
		jobPosition string = "DevOps"
		employeeID int32 = 1222254
	)

	yearsofService := 5
	isFlltime := true
	hourlyRate :=30.4


	weekleyPay := hourlyRate * 3

	fmt.Println(employeeName, jobPosition, employeeID, yearsofService, isFlltime, weekleyPay)



    // Initialize using field names (recommended)
    u1 := User{
        ID:       1,
        Name:     "Alice",
        Email:    "alice@example.com",
        IsActive: true,
    }

    // Initialize without field names (order matters!)
    u2 := User{2, "Bob", "bob@example.com", false}

	fmt.Println(u1)
	fmt.Println(u2)

	    // Access fields
    fmt.Println("Name:", u1.Name)

    // Modify fields
    u1.IsActive = false
    fmt.Println("Updated User:", u1)
}
