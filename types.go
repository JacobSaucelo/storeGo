package main

type StoreType struct {
	Products []StoreItemType `json:store`
}

type StoreCart struct {
	Items []StoreItemType
}

type StoreItemType struct {
	Name  string `json:name`
	Price int16  `json:price`
}
