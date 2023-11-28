package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func mainMenu() {
	clearScreen()
	fmt.Println("Welcome to the Company store. \nUse words BUY and INFO on any item. \nOrder tools in bulk by typing a number.")
	fmt.Println("---------------------------------")
	chooseStore(getStoreItems())
}

func getStoreItems() []StoreItemType {
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

	return itemsContainer
}

func chooseStore(items []StoreItemType) StoreItemType {
	reader := bufio.NewReader(os.Stdin)

	pick, _ := reader.ReadString('\n')

	userInput := strings.ToLower(strings.TrimSpace(pick))

	if contains(userInput, items) {
		fmt.Println("contains")
		return items[userInput]
	}

	fmt.Println("Not a valid input.")

	chooseStore(items)
}

func contains(v string, a []StoreItemType) bool {
	for _, i := range a {
		if i.Name == v {
			return true
		}
	}
	return false
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// "You have requested to order string. Amount: byte. \n Total cost of items: >int32 \n Please CONFIRM or DENY."
