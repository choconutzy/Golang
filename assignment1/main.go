package main

import "fmt"

func main() {
	i, percent, j, russia := 21, `%`, true, `Я`
	num1, num2, num3 := 10, 8, 16
	var k float64 = 123.456
	fmt.Printf("menampilkan nilai i : %d\n", i)
	fmt.Printf("menampilkan tipe data i : %T\n", i)
	fmt.Printf("menampilkan tanda %s\n", percent)
	fmt.Printf("menampilkan tanda %% \n")
	fmt.Printf("menampilkan nilai boolean j %t\n", j)
	j = false
	fmt.Printf("menampilkan nilai boolean j %t\n", j)
	fmt.Printf("menampilkan unicode russia %s\n", russia)
	fmt.Printf("menampilkan nilai base %d %o\n", num1, 17)
	fmt.Printf("menampilkan nilai base %d %o\n", num2, 21)
	fmt.Printf("menampilkan nilai base %d %x\n", num3, 15)
	fmt.Printf("menampilkan nilai base %d %X\n", num3, 15)
	r := rune('Я')
	test := fmt.Sprintf("%U", r)
	fmt.Printf("menampilkan unicode karakter %s\n", test)
	fmt.Printf("menampilkan float %f\n", k)
	fmt.Printf("menampilkan float %e\n", k)
}
