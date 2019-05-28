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

func loop() {
	// I'm thinking the designer of Go must be a foolish

	var a = 0
	for a = 0; a < 10; a++ {
		fmt.Println(a)
	}

	for b := 0; b < 10; b++ {
		fmt.Println(b)
	}
}

func main() {
	varSometing()
}
