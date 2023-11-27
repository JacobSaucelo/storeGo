package main

import (
	"fmt"
	"os"
)

func mainMenu() {
	clearScreen()
	fmt.Println("Welcome to the Company store. \nUse words BUY and INFO on any item. \nOrder tools in bulk by typing a number.")
	fmt.Println("---------------------------------")
}

func getStoreItems() {
	jsonFile, err := os.Open("./assets/store.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	fmt.Printf("file %T", jsonFile)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
