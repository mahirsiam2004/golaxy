package golang
package main

import (
    "fmt"
    "strings"
)

func main() {
    
    message := "Hello, Golang!"

  
    fmt.Println("Original string:", message)

    fmt.Println("Length:", len(message))

 
    fmt.Println("First character:", string(message[0]))
    fmt.Println("Last character:", string(message[len(message)-1]))

    fmt.Println("Uppercase:", strings.ToUpper(message))
    fmt.Println("Lowercase:", strings.ToLower(message))

    fmt.Println("Contains 'Go':", strings.Contains(message, "Go"))


    replaced := strings.Replace(message, "Golang", "World", 1)
    fmt.Println("After replace:", replaced)

 
    parts := strings.Split(message, " ")
    fmt.Println("Split by space:", parts)

 
    joined := strings.Join(parts, "-")
    fmt.Println("Joined with '-':", joined)
}
