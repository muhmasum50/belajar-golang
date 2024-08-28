package main

import (
	"fmt"
)

func main() {
	//var person map[string]string = map[string]string{}
	//
	//person["name"] = "Bang Pixel"
	//person["address"] = "Sidoarjo"

	person := map[string]string{
		"name":    "Bang Pixel",
		"address": "Sidoarjo",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := map[string]string{}
	book["title"] = "Buku Ajaib"
	book["author"] = "Bang Pixel"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
