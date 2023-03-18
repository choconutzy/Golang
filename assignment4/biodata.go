package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {
	input := (os.Args[1])
	// Convert the string to an integer using strconv.Atoi()
	num, err := strconv.Atoi(input)
	// Check for errors during the conversion
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	num--
	friends := []Friend{
		{Name: "Andi", Address: "Jl. Merdeka No. 10", Job: "Student", Reason: "Untuk keperluan tugas akhir"},
		{Name: "Budi", Address: "Jl. Sudirman No. 15", Job: "Backend Developer", Reason: "Upgrade skill"},
		{Name: "Cindy", Address: "Jl. Hayam Wuruk No. 30", Job: "Software Developer", Reason: "Upgrade Skill"},
	}
	fmt.Printf("Nama 		: %s\n", friends[num].Name)
	fmt.Printf("Alamat 		: %s\n", friends[num].Address)
	fmt.Printf("Pekerjaan 	: %s\n", friends[num].Job)
	fmt.Printf("Alasan 		: %s\n", friends[num].Reason)
}
