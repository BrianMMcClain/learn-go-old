package main

import "fmt"

func isBig(x int) bool {
	if x < 100 {
		return false
	} else {
		return true
	}
	
	return false
}

func main() {
	fmt.Println(isBig(99))
	fmt.Println(isBig(101))
}