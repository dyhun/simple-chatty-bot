package main

import "fmt"

func f(x int) int {
    if x < 0 {
        return f1(x)
    }
    return f2(x)
}


func f1(x int) int {
    return 2*x
}


func f2(x int) int {
    return 3*x
}