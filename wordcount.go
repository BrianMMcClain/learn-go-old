package main

import (
	"tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	for _, v := range strings.Fields(s) {
		m[v] += 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}