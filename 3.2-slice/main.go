package main

import (
	"fmt"
	"strings"
)

func main() {

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)

	fmt.Println(names[0]) // Alice

	names[1] = "David"
	fmt.Println(names) // ["Alice", "David", "Charlie"]

	// Append elements
	names = append(names, "Emma")
	fmt.Println(names)

	// 4. Loop over a slice
	fmt.Println("//------------------ 4. Loop over a slice")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i]) 
	}

	for i, name := range names {
    fmt.Printf("%d: %s\n", i, name)
	}
	fmt.Println("//------------------")
	// 5. Pass slice to function (modifies original)
	fmt.Println("5. Pass slice to function (modifies original)")
	toUpper(names)
	fmt.Println(names)

	fmt.Println("//------------------")
	// 6. Create slice with make()
	fmt.Println("6. Create slice with make()")
	numbers := make([]int, 5)        // length 5
	numbersCap := make([]int, 5, 10) // length 5, capacity 10


	fmt.Println(numbers)
	fmt.Println(numbersCap)

	fmt.Println("//------------------")
	fmt.Println("7. Slicing a slice")
	nums := []int{10, 20, 30, 40, 50}

	part := nums[1:4]
	fmt.Println(part) // [20 30 40]

	fmt.Println("//------------------")


	//8. Remove element (index 2)
	fmt.Println("//8. Remove element (index 2)")
	i := 2
	nums = append(nums[:i], nums[i+1:]...)

	fmt.Println("//------------------")


	//9. Check length & capacity
	fmt.Println("/9. Check length & capacity")
	fmt.Println(len(nums))
	fmt.Println(cap(nums))


	fmt.Println("//------------------")


}

func toUpper(names []string) {
    for i := range names {
        names[i] = strings.ToUpper(names[i])
    }
}