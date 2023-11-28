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

	for {
		var playerInput string
		fmt.Scan(&playerInput)

		if strings.ToLower(playerInput) == "exit" {
			break
		}
		if cart.BuyItem(playerInput, store) {
			cart.ShowCartItems()
		}
	}

	fmt.Println("You bought:")
	cart.ShowCartItems()

}

// go run main.go methods.go generate.go types.go
