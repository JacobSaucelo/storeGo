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

func GetStoreItems() []StoreItemType {
	jsonFile, err := os.Open("./assets/store.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	storeByteValue, _ := ioutil.ReadAll(jsonFile)

	var itemsContainer []StoreItemType

	json.Unmarshal(storeByteValue, &itemsContainer)

	return itemsContainer
}

func (items *[]StoreItemType) DisplayStore() {
	for _, item := range items {
		fmt.Printf("*%s // >%d", item.Name, item.Price)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// "You have requested to order string. Amount: byte. \n Total cost of items: >int32 \n Please CONFIRM or DENY."
// - shouldve used map for more optimized
