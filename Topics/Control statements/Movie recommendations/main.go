package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	if age <= 14 {
		fmt.Println("Toy Story 4")
	} else if age <= 18 {
		fmt.Println("The Matrix")
	} else if age <= 25 {
		fmt.Println("John Wick")
	} else if age <= 35 {
		fmt.Println("Constantine")
	} else if age > 35 {
		fmt.Println("Speed")
	}
}
