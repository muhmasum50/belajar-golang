package main

import "fmt"

func main() {
	type NomorKTP string

	var KtpBangPixel NomorKTP = "22222222222222"

	var contoh string = "11111111111111"
	var contohKTP NomorKTP = NomorKTP(contoh)

	for i := 0; i < 10; i++ {

	}

	fmt.Println(KtpBangPixel)
	fmt.Println(contohKTP)
}
