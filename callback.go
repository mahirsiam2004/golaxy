package main
import "fmt"

func process(name string, callback func(string)) {
    fmt.Println("Processing...")
    callback(name)
}

func greet(n string) {
    fmt.Println("Hello,", n)
}

func main() {
    process("Mahir", greet)
}
