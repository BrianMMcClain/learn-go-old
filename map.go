package main

import "fmt"

func main() {
	m := make(map[string]bool)
	m["apple"] = true
	m["orange"] = true
	m["banana"] = false
	
	fmt.Println(m)
	fmt.Println(m["banana"])
}