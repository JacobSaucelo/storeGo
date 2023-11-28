// playing with creating json files

package main

import (
	"fmt"
	"os"
	"strconv"
)

func createString() string {
	texts := `{"Products":[`

	for _, value := range storeItems {
		texts += `{"Name":"` + value.Name + `","Price":` + strconv.Itoa(int(value.Price)) + "},\n"
	}

	texts += "]}"

	return texts
}

func generate() {
	g := []byte(createString())
	folderErr := os.MkdirAll("assets", os.ModePerm)
	if folderErr != nil {
		fmt.Println("cant create folder.")
	}
	fileErr := os.WriteFile("assets/store.json", g, 0644)
	if fileErr != nil {
		panic(fileErr)
	}
	fmt.Println("Successfully generated store.json")
}

var storeItems = []StoreItemType{
	{
		Name:  "Walkie-talkie",
		Price: 12,
	},
	{
		Name:  "Flashlight",
		Price: 15,
	},
	{
		Name:  "Shovel",
		Price: 30,
	},
	{
		Name:  "Lockpicker",
		Price: 20,
	},
	{
		Name:  "Pro-flashlight",
		Price: 25,
	},
	{
		Name:  "Stun grenade",
		Price: 40,
	},
	{
		Name:  "Boombox",
		Price: 60,
	},
	{
		Name:  "TZP-Inhalant",
		Price: 120,
	},
	{
		Name:  "Zap gun",
		Price: 400,
	},
	{
		Name:  "Jetpack",
		Price: 700,
	},
	{
		Name:  "Extension ladder",
		Price: 60,
	},
	{
		Name:  "Radar-booster",
		Price: 50,
	},
}
