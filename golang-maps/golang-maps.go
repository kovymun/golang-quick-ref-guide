/*
-------
GO MAPS
-------
*/
package main

import "fmt"

func main() {
	// Defining a studentIds map
	studentIds := map[string]int{
		"James Malcom": 12981,
		"Judy Jetson":  98723,
	}

	// Printing studentIds map to the console
	fmt.Println(studentIds)

	// Retrieve and print an ID value
	fmt.Println("Judy's ID number:", studentIds["Judy Jetson"])
}

/*
--------
DETAILS:
--------
1. What is a Map in Go? It is a built-in data structure that stores data as key-value pairs, enabling efficient data retrieval based on unique keys.
If you're familiar with other programming languages, a map in Go is similar to an object in JavaScript or a dictionary in Python.

2. SYNTAX:
    mapVariableName := map[keyType]valueType{
    key1: value1,
    key2: value2,
}

EXAMPLE:

*/
