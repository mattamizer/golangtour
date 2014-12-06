package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordmap := make(map[string]int)
	for _, index := range strings.Fields(s) {
		wordmap[index]++
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
