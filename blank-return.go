package main

import "fmt"

func getSomeStrings() (a, b string) {
	a = "Hello"
	b = "World"
	return
}

func main() {
	a, b := getSomeStrings()
	fmt.Println(a, b)
}