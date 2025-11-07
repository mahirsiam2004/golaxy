package golang
package main

import "fmt"

func main() {
    // Declare an array of 5 integers
    var numbers [5]int

    // Assign values to the array
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    // Print the entire array
    fmt.Println("Array:", numbers)

    // Access individual elements
    fmt.Println("First element:", numbers[0])
    fmt.Println("Last element:", numbers[4])

    // Loop through the array
    fmt.Println("Array elements:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("Index %d: %d\n", i, numbers[i])
    }

    // Calculate the sum of array elements
    sum := 0
    for _, value := range numbers {
        sum += value
    }
    fmt.Println("Sum of all elements:", sum)
}
