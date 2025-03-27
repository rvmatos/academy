package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		showIntroduction()
		showMenu()

		option := readOption()

		switch option {
		case 1:
			startMonitoring()
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
}

func showIntroduction() {
	var name = "Reinaldo"
	var version float32 = 1.2

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

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.rvmatos.com"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded successfuly")
	} else {
		fmt.Println("Site:", site, "failed to load. Status Code:", response.StatusCode)
	}
}
