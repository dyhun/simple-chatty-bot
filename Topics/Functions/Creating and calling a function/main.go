package main

import "fmt"

func message(name string) string {
	return name + " is learning how to call functions!"
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)
	fmt.Println(message(userName))
}
