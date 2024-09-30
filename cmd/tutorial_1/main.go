package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	fmt.Println("Hello World")

	fmt.Println("Storing Float")

	var floatNum float64 = 123.5
	fmt.Println(floatNum)

	var floatBit float32 = 1324.99
	fmt.Println(floatBit)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 45
	var sum = floatNum32 + float32(intNum32)
	fmt.Println(sum)

	var intNum1 int = 3
	var intNum2 int = 4
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var lineBreak string = "Line \nBreak"
	fmt.Println(lineBreak)

	var myString string = "Four"
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)
	
	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3) // Default value of int is 0
	// Default value of boolean is false
	// Default value of string is ''

	var testVar = "Printing da string"
	fmt.Println(testVar)

	myVar := "print da string"
	fmt.Println(myVar)

	var1, var2 := 1,2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}



