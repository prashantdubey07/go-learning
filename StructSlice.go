// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
    // Your code goes here.
    // Create a slice of Color objects
    colorSlice := make([]Color, 0, len(colorNames))
    for index, value := range colorNames {
		tempColor := Color{value, hexValues[index]}
        colorSlice = append(colorSlice,tempColor)
	}
    return colorSlice
}

type Color struct {
	Name string
	Hex  int
}
