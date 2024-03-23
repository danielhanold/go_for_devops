package say

import "fmt"

// Public function.
func PrintHello() {
	fmt.Println("Hello")
}

// Private function.
func printWorld() {
	fmt.Println("World")
}

// Public function.
func PrintHelloWorld() {
	PrintHello()
	printWorld()
}
