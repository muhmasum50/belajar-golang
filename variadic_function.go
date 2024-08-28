package main

import "fmt"

func main() {
	resultSumAll := sumAll(10, 10, 10)

	fmt.Println("Result : ", resultSumAll)

	/** jika parameter berupa slice */
	numbers := []int{10, 10, 10, 10, 10, 10}
	total := sumAll(numbers...)
	fmt.Println("Result : ", total)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
