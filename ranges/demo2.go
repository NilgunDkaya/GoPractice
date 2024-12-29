package ranges

import "fmt"

func Demo2() {
	/*var toplam int = 0

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			toplam = toplam + i
		}
	}
	fmt.Println(toplam) */

	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			toplam = toplam + sayi
		}
	}
	fmt.Println("Toplam: ", toplam)

}
