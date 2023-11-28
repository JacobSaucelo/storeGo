package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mainMenu()
	store := GetStoreItems()
	cart := StoreCart{}
	store.DisplayStore()

	for {
		reader := bufio.NewReader(os.Stdin)

		playerInput, _ := reader.ReadString('\n')

		fmt.Println("DEBUG:", strings.TrimSpace(playerInput))

		if strings.ToLower(playerInput) == "exit" {
			break
		}
		if cart.BuyItem(strings.TrimSpace(playerInput), store) {
			cart.ShowCartItems()
		}
	}

	fmt.Println("You bought:")
	cart.ShowCartItems()

}

// TODO  add concatenate shortcut "buy <item>"
// go run main.go methods.go generate.go types.go
