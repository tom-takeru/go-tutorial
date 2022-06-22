package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	var wordCountMap = make(map[string]int)
	for _, v := range words {
		wordCountMap[v]++
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
