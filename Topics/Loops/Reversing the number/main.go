package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	strNumber := strconv.Itoa(number)
	reverseStrNumber := ""
	for length := len(strNumber); length > 0; length-- {
		reverseStrNumber += string(strNumber[length-1])
	}
	strconv.Atoi(reverseStrNumber)
	fmt.Println(reverseStrNumber)
}
