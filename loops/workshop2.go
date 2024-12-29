package loops

import "fmt"

func Demo4() {

	number := 0

	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&number)

	isNumberPrime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isNumberPrime = false
		}
	}

	if isNumberPrime {
		fmt.Println("Tebrikler! Girdiğiniz sayı asal")
	} else {
		fmt.Println("Girdiğiniz sayı asal değildir")
	}

}
