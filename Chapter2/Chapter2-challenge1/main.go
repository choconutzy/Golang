package main

import "Chapter2-challenge1/routers"

func main() {

	PORT := ":4000"
	routers.StartServer().Run(PORT)
}
