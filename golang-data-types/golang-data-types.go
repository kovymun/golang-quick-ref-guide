package main

import "fmt"

func main() {
	fName := "James" //inferred as type string
	var age int = 10 //string type integer
	fmt.Println(fName, age)
}

/*
--------------
GO DATA-TYPES:
--------------
* DATA-TYPES USED IN GO: string | int | bool | float64 | array
* SYNTAX: var *identifier-name* *data-type* = initialized value
* EXAMPLE: var fName string = "James"
* IMPORTANT: When you declare a variable using the := syntax, Go automatically infers the type based on the assignment value.
*/
