package main

import "fmt"

func main() {
	x := 10
	switch x {
		case 5:
			fmt.Println("Five")
		case 10:
			fmt.Println("Ten")
		case 20:
			fmt.Println("Twenty")
		default:
			fmt.Println("I don't even know what you did")
	}
}