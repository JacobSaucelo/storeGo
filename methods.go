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
	fmt.Println("Welcome to the Company store. \nUse words BUY or EXIT and pick any item. \nOrder tools in bulk by typing a number.")
	fmt.Println("---------------------------------")
}

func MainMenuOptions() {
	reader := bufio.NewReader(os.Stdin)

	menuPicked, _ := reader.ReadString('\n')

	switch strings.TrimSpace(menuPicked) {
	case "buy":
		fmt.Println("Buy")
	case "info":
		fmt.Println("info")
	case "exit":
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid input...")
		MainMenuOptions()
	}
}

func GetStoreItems() StoreType {
	jsonFile, err := os.Open("./assets/store.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	storeByteValue, _ := ioutil.ReadAll(jsonFile)

	var itemsContainer StoreType

	json.Unmarshal(storeByteValue, &itemsContainer)

	return itemsContainer
}

func (items *StoreType) DisplayStore() {
	for _, item := range items.Products {
		fmt.Printf("*%s // >%d \n", item.Name, item.Price)
	}
}

func (items *StoreType) BuyItem(itemName string, store StoreType) bool {
	for _, item := range store.Products {
		if strings.EqualFold(item.Name, itemName) {
			items.Products = append(items.Products, item)
			fmt.Printf("%s added to your cart.\n\n", item.Name)
			return true
		}
	}
	fmt.Printf("'%s' not found.", itemName)
	return false
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// "You have requested to order string. Amount: byte. \n Total cost of items: >int32 \n Please CONFIRM or DENY."
// - shouldve used map for more optimized
