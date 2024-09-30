package main

import (
	"errors"
	"fmt"
)

func main(){
	var printValue string = "Passing the value"
	printMe(printValue)

	var numerator int = 55
	var denominator int = 9
	var result, remainder, err = intDivision(numerator, denominator)

	// If statement
	if err!=nil{
		fmt.Printf(err.Error())
	} else if remainder==0{
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer is %v with remainder %v", result, remainder)
	}

	// Switch statement
	// No need for Break in golang
	switch {
	case err!=nil:
		fmt.Printf(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division is %v", result)
	case remainder!=0:
		fmt.Printf("The result of the integer division is %v and the remainder is %v", result, remainder)

	}


	// Switch with a conditional statement
	switch remainder{
	case 0:
		fmt.Printf("The division was exact")
	case 1,2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	// Default value of error is nil
	var err error 
	if denominator == 0{
		err = errors.New("Cannot Divide by Zero")
		return 0,0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}