package main

import "fmt"

func main() {
	sayWhat()
	sayWhatTo("Bang Pixel")

	fmt.Println("-----------------------------------------------")

	resultGetHello := getHello("Bang Pixel")
	fmt.Println(resultGetHello)

	fmt.Println("-----------------------------------------------")

	firstName, _ := getFullName()
	fmt.Println(firstName)

	fmt.Println("-----------------------------------------------")

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func sayWhat() {
	fmt.Println("Say Whaaattt")
}

func sayWhatTo(name string) {
	fmt.Println("Helooo, "+name, ". Have a nice day")
}

// return value
func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

// return multiple values
func getFullName() (string, string) {
	return "Bang", "Pixel"
}

// return multiple named values
func getCompleteName() (firstName string, lastName string, title string) {
	firstName = "Bang"
	lastName = "Pixel"
	title = "S.Kom"

	return firstName, lastName, title
}
