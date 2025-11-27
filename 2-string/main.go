package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "iman"
	age := 25
	description := "Hello this is my go learning go code"
	println("hello world, my name is: ", name , "and my age is : ", age)
	println("hello")

	fmt.Println(strings.Contains(description, "code"))
	fmt.Println(strings.Contains(description, "golan"))

	fmt.Println(strings.Count(description, "go"))


	mystringArray := strings.Split(description, " ")
	fmt.Println(mystringArray)

	fmt.Println("the len of string is : ", len(mystringArray))

	for idx, item := range mystringArray {
		println("idx is : ", idx, " and item is",item)
	}

	stringCommaSep := strings.Join(mystringArray, ",")
	fmt.Println(stringCommaSep)

	println(strings.Repeat("iman ", 10))

	println("-----------------compare------------------")
	println(strings.Compare("golang", "golang"))   // if return 0 , it mean the strings are equal
	println(strings.Compare("golang", "GOLANG"))	// if return non-zero , it mean the strings aren't equal
	println(strings.EqualFold("golang", "GOLANG"))  // compare but case-insensitive
}