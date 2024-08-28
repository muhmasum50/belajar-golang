package main

import "fmt"

func main() {
	goodbye := getGoodBye

	fmt.Println(goodbye("Bang Pixel "))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
