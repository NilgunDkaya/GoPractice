package structs

import "fmt"

func Demo1() {
	//fmt.Println(product{"Bilgisayar", 500, "XYZ", 0})
	fmt.Println(product{name: "Bilgisayar", price: 500, brand: "XYZ"})
	fmt.Println(product{"Telefon", 800, "ABC", 0})
}

type product struct {
	name         string
	price        float64
	brand        string
	discountRate int
}
