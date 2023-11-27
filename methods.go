package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	storeByteValue, _ := ioutil.ReadAll(jsonFile)

	var itemsContainer []StoreItemType

	json.Unmarshal(storeByteValue, &itemsContainer)

	for _, item := range itemsContainer {
		fmt.Println("*", item.Name, " // >", item.Price)
	}

	fmt.Println("items: ", itemsContainer)

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
