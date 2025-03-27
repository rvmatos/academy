package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()

	option := readOption()

	switch option {
	case 1:
		fmt.Println("Monitoring...")

	case 2:
		fmt.Println("Showing logs...")

	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("Invalid Option.")
		os.Exit(-1)
	}
}

func showIntroduction() {
	var name = "Reinaldo"
	var version float32 = 1.1

	fmt.Println("Hello, mr.", name)
	fmt.Println("This program is running in the version", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}
