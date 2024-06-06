package main

import "fmt"

/*
Issues and Recommendations:

1. Factorial Calculation Error: The calculation for factorial is incorrect. 
The code is summing up the integers instead of multiplying them. 
Additionally, the result variable should be an integer, not a float.
    
2. Function Naming Convention: The printfactorial function should follow camel case, i.e., printFactorial.
    
3. Return Type Consistency: The function factorial should return an integer, but the return type is incorrect (float64).
*/

// Function to calculate the factorial of a number
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Function to print the factorial of a number
func printFactorial() {
	num := 5
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}

func main() {
	printFactorial()
}

