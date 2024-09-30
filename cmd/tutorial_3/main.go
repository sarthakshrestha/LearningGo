package main

import "fmt"
func main(){


	// Can be initialized with the type of Array
	// Like this

	// var intArr1 [3]int32 = [3]int32{1,2,3}

	// intArr := [3]int32{1,2,3}

	intArr := [...]int32{1,2,3} // Can be declared with 
	// an array size with dot dots as well.


	// Fixed Length of 3
	// Elements are of type int32
	// Access element of 1,2,
	intArr[1] = 21
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// Print location of the memory using ampersand 
	// symbol
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices = Wrappers around array
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}

