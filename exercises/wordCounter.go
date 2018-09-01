package main

import (
  	"../goLang/wc"
  	"strings"
)

func WordCounter(s string) map[string]int {
  	words := strings.Fields(s)
  	count := make(map[string]int)

  	for i := 0; i < len(words); i++ {
    	count[words[i]] += 1
  	}

  	return count
}

func main() {
  	wc.Test(WordCounter)
}
