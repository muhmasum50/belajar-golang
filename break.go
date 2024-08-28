package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		if i == 2 {
			break
		}

		fmt.Println("Loop Break ke", i)
	}

	// continue
	for idx := 1; idx < 10; idx++ {
		if idx%2 == 0 {
			continue
		}

		fmt.Println("Loop Continue ke", idx)
	}
}
