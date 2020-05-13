package main

import "fmt"

func main() {

	// First form (i is block scope)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----")

	// Second form
	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 2
	}

	// Nested loop with break
	i = 0
	for i < 4 {
		for j := 0; j < 10; j++ {
			if j > 5 {
				fmt.Println("Breaking...")
				break
			}
			fmt.Println("Inner J=", j)
		}
		fmt.Println("Outer I=", i)
		i++
	}

}
