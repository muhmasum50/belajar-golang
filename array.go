package main

import "fmt"

// data
func main() {

	var names [3]string

	names[0] = "Pixel"
	names[1] = "Gaming"
	names[2] = "YT"

	fmt.Println(names)

	var values = [...]int{
		90,
		100,
		80,
		54,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 125
	fmt.Println(values)
}
