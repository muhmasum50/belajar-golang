package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Coba looping nih : ", counter)
		counter++
	}

	fmt.Println("-----------------------------------------------")

	for count := 1; count <= 10; count++ {
		fmt.Println("coba looping nih : ", count)
	}

	fmt.Println("-----------------------------------------------")

	datas := []string{"Halo", "Test", "Coba", "Array"}

	for i := 0; i < len(datas); i++ {
		fmt.Println(datas[i])
	}

	fmt.Println("-----------------------------------------------")

	for index, data := range datas {
		fmt.Println(data, index)
	}

	fmt.Println("-----------------------------------------------")

	/** for without index */
	for _, data := range datas {
		fmt.Println(data)
	}

}
