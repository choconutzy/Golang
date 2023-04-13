package main

import "fmt"

func loop(x string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai %s = %d\n", x, i)
	}
}

func loopUnicode(y string) {
	for i := 0; i <= 12; i += 2 {
		r := rune(y[i])

		fmt.Printf("charakter %U '%s' start at byte position %d\n", r, string(y[i]), i)
	}
}

func main() {
	init := "i"
	for i := 0; i <= 10; i++ {
		if i < 5 && init == "i" {
			loop(init)
			init = "j"
			i = 0
		}
		if i < 5 && init == "j" {
			fmt.Printf("Nilai %s = %d\n", init, i)
		}
		if i == 5 {
			loopUnicode("CTAMLDASPDBUO")
		}
		if i > 5 {
			fmt.Printf("Nilai %s = %d\n", init, i)
		}
	}
}
