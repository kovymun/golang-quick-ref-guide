/*
---------------------
GO ARRAYS ANS SLICES:
---------------------
*/

package main

import "fmt"

func main() {
	// Printing output to the console.
	fmt.Println(demoArrays())
	fmt.Println(demoSlices())
}

// Syntax for Go function demonstrating arrays (ex: Motor vehicles)
func demoArrays() string {
	// Declaring and initializing an array
	var cars [3]string
	cars[0] = "Toyota"
	cars[1] = "Honda"
	cars[2] = "Ford"

	// Accessing and displaying array elements
	output := "Array elements: "
	for _, car := range cars {
		output += car + " "
	}
	return output
}

// Syntax for Go function demonstrating slices
func demoSlices() string {
	// Declaring and initializing a slice
	vegetables := []string{"Carrot", "Potato", "Cucumber"}

	// Adding elements to a slice
	vegetables = append(vegetables, "Spinach", "Tomato")

	// Accessing and displaying slice elements
	output := "Slice elements: "
	for _, vegetable := range vegetables {
		output += vegetable + " "
	}
	return output
}

/*
-----------------------------
GO ARRAYS ANS SLICES DETAILS:
-----------------------------
* What is an Array?
  - An array is a fixed-size collection of elements of the same data type.
  - SYNTAX:
    var arrayName [size]dataType
    Example: var numbers [5]int

* What is a Slice?
  - A slice is a dynamically-sized, flexible view of an array.
  - Unlike arrays, slices can grow or shrink as needed.
  - SYNTAX:
    sliceName := []dataType{values...}
    Example: fruits := []string{"Apple", "Banana"}

* Key Differences Between Arrays and Slices:
  - Arrays have a fixed size; slices are dynamic.
  - Slices can use built-in functions like append() to add elements.
  - Slices are more commonly used in Go due to their flexibility.

* Useful Functions for Slices:
  - len(slice) - Returns the number of elements in the slice.
  - cap(slice) - Returns the capacity of the slice.
  - append(slice, element) - Adds one or more elements to the slice.

* Why use Arrays and Slices?
  - Use arrays for fixed collections of data.
  - Use slices for dynamic, flexible data structures.
*/
