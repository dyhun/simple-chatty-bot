package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	sum := 1
	for i := 1; i <= number; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
