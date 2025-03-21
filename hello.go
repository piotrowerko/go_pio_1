package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    sum := add(3, 5)
    fmt.Println("Sum of 3 and 5 is:", sum)

    // Create an instance of the Calculator struct
    calc := Calculator{}
    result := calc.Multiply(4, 5)
    fmt.Println("Product of 4 and 5 is:", result)

    // Use a function from the utils file
    PrintMessage("This is a message from the utils file")
}

// add takes two integers and returns their sum
func add(a int, b int) int {
    return a + b
}

// Calculator is a struct with methods
type Calculator struct{}

// Multiply takes two integers and returns their product
func (c Calculator) Multiply(a int, b int) int {
    return a * b
}