package main

import "fmt"

// Function to add two numbers
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// Function to return sum and multiplication
func getNumbers(num1 int, num2 int) (int, int) {
	sum2 := num1 + num2
	mul := num1 * num2
	return sum2, mul
}

func main() {
	a := 10
	b := 20

	ans := add(a, b)
	fmt.Println("Addition:", ans)

	p, q := getNumbers(a, b)
	fmt.Println("Sum:", p)
	fmt.Println("Multiplication:", q)
}
