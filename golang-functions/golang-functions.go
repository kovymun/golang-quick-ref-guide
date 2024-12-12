package main

import "fmt"

func main() {
	//printing output to the console.
	fmt.Println(names())
}

// Syntax for a go lang function
func names() string {
	return "This is my name"
}

/*
------
NOTES:
------

* What is a function? A function is small reusable piece of code that contains a specific action/logic.
* SYNTAX:
	function functionName() data-type(to be returned from the function) {
		//function body logic
	}
*/
