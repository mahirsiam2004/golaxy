package main

import "fmt"

func dummy(){
	test:=func(){
		fmt.Println("Hello, World!")
	}
	test()
}


func main() {
	dummy()
}