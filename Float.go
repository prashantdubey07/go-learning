// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
 import (
 	//"fmt"
 	"strconv"
 )

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the the result of adding 2 numbers 
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
    // Your code goes here.

	// Convert the first string to a float64
	floatValue1, _ := strconv.ParseFloat(value1, 64)
	/*
	if err != nil {
		fmt.Println("Error parsing float64:", err)
	} else {
		fmt.Printf("Parsed float64: %v, Type: %T\n", floatValue1, floatValue1)
	}
	*/
	// Convert the second string to a float64

	floatValue2, _ := strconv.ParseFloat(value2, 64)
	/*if err != nil {
		fmt.Println("Error parsing float64:", err)
	} else {
		fmt.Printf("Parsed float64: %v, Type: %T\n", floatValue2, floatValue2)
	}*/
	// Calculate and return the result

	return floatValue1+floatValue2
}
