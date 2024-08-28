package main

import "fmt"

func main() {
	name := "bang piexl"

	if name == "pixel" {
		fmt.Println("hello pixel")
	} else if name == "bang" {
		fmt.Println("hello bang")
	} else {
		fmt.Println("hi, nilai disini false")
	}

	if length := len(name); length > 5 {
		fmt.Println("name is too long")
	} else {
		fmt.Println("name is correct")
	}

	switch name {
	case "pixel":
		fmt.Println("pixel")
	default:
		fmt.Println("ya oke ")
	}
}
