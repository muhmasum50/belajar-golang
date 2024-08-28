package main

import "fmt"

func main() {

	sayHelloWithFilter("Bang Pixel", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "***"
	}

	return name
}
