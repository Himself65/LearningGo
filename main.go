package main

import "fmt"

func printSomething() {
	fmt.Println("Hello, world")
}

func varSometing() {
	a := 1
	fmt.Println(a)

	b := "Stirng"
	fmt.Println(a, b)

	// mark: '\n' is different from "\n"
	fmt.Println()
	fmt.Print(a, b, "\n")

	fmt.Println()
	fmt.Print(a, b, '\n')
}

func main() {
	varSometing()
}
