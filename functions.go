package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mult(x int, y int) int {
	return x * y
}

func div(x int, y int) int {
	return x / y
}
	
	
func main() {
	fmt.Println(add(1,2))
	fmt.Println(sub(1,2))
	fmt.Println(mult(1,2))
	fmt.Println(div(1,2))
}