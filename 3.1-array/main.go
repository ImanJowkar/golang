package main

import (
	"fmt"
	"strings"
)


func main() {
	//array
	fmt.Println("------------------- Array --------------------------")
	var myary1 [2]int
	var myary2 [8]int = [8]int {1,2}

	myary3 := [14]int{1,2}
	myary4 := [...]float32{1.2,1.5,1.9}  // dynamic len size detection

	myary5 := []string{    // slice , If you want a dynamic list instead, use a slice:
		"Alice", "Bob", "Charlie", "David", "Emma",
		"Frank", "Grace", "Helen", "Ivan", "Julia",
	}

	fmt.Println(myary1)
	fmt.Println(myary2)
	fmt.Println(myary3)
	fmt.Println(myary4)
	fmt.Println(myary5)

	// len of an array

	fmt.Println("len array1:", len(myary1))
	fmt.Println("len array2:", len(myary2))
	fmt.Println("len array3:", len(myary3))
	fmt.Println("len array4:", len(myary4))


	// change the value of an array
	myary1[0] = 88
	myary2[5] = -55
	myary3[10] = 43333221
	myary4[2] = -3333333.2
	fmt.Println(myary1)
	fmt.Println(myary2)
	fmt.Println(myary3)
	fmt.Println(myary4)

    var names [10]string = [10]string{
        "Alice",
        "Bob",
        "Charlie",
        "David",
        "Emma",
        "Frank",
        "Grace",
        "Helen",
        "Ivan",
        "Julia",
    }

	searchKey := "Frank"
	for idx, name := range names{
		if name == searchKey {
			fmt.Println("the index of ", name , "is: ", idx)
		}
	}


    fmt.Println("Before:", names)

    // Pass array by pointer
    toUppercase(&names)

    fmt.Println("After:", names)
}


func toUppercase(names *[10]string) {
    for i := 0; i < len(names); i++ {
        names[i] = strings.ToUpper(names[i])
    }
}