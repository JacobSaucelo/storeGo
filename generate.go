// playing with creating json files

package main

import "strconv"

type storeType struct {
	store []itemType
}

type itemType struct {
	name  string
	price int16
}

func createString() string {
	texts := "["

	for _, value := range storeItems {
		texts += "{'name':" + value.name + ",'price':" + strconv.Itoa(int(value.price)) + "},\n"
	}

	texts += "]"

	return texts
}

func generate(texts string) {
	g := []byte(texts)
}

var storeItems = []itemType{
	{
		name:  "Walkie-talkie",
		price: 12,
	},
	{
		name:  "Flashlight",
		price: 15,
	},
	{
		name:  "Shovel",
		price: 30,
	},
	{
		name:  "Lockpicker",
		price: 20,
	},
	{
		name:  "Pro-flashlight",
		price: 25,
	},
	{
		name:  "Stun grenade",
		price: 40,
	},
	{
		name:  "Boombox",
		price: 60,
	},
	{
		name:  "TZP-Inhalant",
		price: 120,
	},
	{
		name:  "Zap gun",
		price: 400,
	},
	{
		name:  "Jetpack",
		price: 700,
	},
	{
		name:  "Extension ladder",
		price: 60,
	},
	{
		name:  "Radar-booster",
		price: 50,
	},
}
