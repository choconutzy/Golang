package main

import (
	"fmt"
	"strings"
)

func main() {
	init := strings.Split("Selamat malam", "")
	result := map[string]int{}
	for i := 0; i < len(init); i++ {
		fmt.Printf(init[i] + "\n")
		if result[init[i]] == 0 {
			result[init[i]] = 1
		} else {
			result[init[i]]++
		}
	}
	fmt.Print(result)
}
