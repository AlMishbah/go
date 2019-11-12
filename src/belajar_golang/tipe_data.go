package main

import "fmt"

func main()  {
	var postiveNumber uint8 = 89
	var negativeNumber = "-1223445677"
	var decimalNumber = 2.28
	var exist bool = true

	fmt.Printf("Bilangan positif : %d\n", postiveNumber)
	fmt.Printf("Bilangan negatif : %d\n", negativeNumber)

	fmt.Printf("Bilangan desimal : %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal : %.5f\n", decimalNumber)

	fmt.Printf("exist ? %t \n", exist)

}