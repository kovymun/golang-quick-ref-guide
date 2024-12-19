package main

import "fmt"

// Syntax for a go lang function
func sum(a int, b int) int {
	return a + b
}

func main() {
	//printing output to the console.
	fmt.Println(sum(1, 6))
}

/*
------------
GO FUNCTIONS
------------

* What is a function? A function is small reusable piece of code that contains a specific action/logic.
* SYNTAX:
	function functionName() data-type(to be returned from the function) {
		//function body logic
	}

*/
