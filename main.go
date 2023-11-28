package main

import (
	"fmt"
	"strings"
)

func main() {
	mainMenu()
	store := GetStoreItems()
	cart := StoreCart{}
	store.DisplayStore()
	MainMenuOptions()

	for {
		var playerInput string
		fmt.Scan(&playerInput)

		if strings.ToLower(playerInput) == "exit" {
			break
		}
		if cart.BuyItem(playerInput, store) {

		}

	}

}

// go run main.go methods.go generate.go types.go
