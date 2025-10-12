package main
import "fmt"

// Higher-order function: takes another function as argument
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

// Function to add two numbers
func add(x int, y int) {
	z := x + y
	fmt.Println(z) 
}

func call() func (x int, y int){
	return add
}

func main() {
	processOperation(2, 5, add)
	sum:=call()
	sum(4,3)
}
