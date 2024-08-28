package main

import "fmt"

func main() {
	names := [...]string{
		"Bang",
		"Pixel",
		"Gaming",
		"YT",
		"Series",
		"Pewdiepie",
	}

	slice := names[4:6]
	fmt.Println(slice)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println("-----------------------------------------------")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	daySlice2 := append(daySlice1, "Ahad")
	fmt.Println(daySlice2)
	fmt.Println("-----------------------------------------------")

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Bang"
	newSlice[1] = "Pixel"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Gaming YT")
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println("-----------------------------------------------")

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// the different between array and slice in go
	thisIsArray := [...]int{1, 2, 3, 4, 5, 6}
	thisIsSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
