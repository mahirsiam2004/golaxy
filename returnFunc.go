package main
import "fmt"

var a = 10 

func call() int {
	return a
}

func main() {
	fmt.Println(call()) 
}
