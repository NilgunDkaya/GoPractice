package interfaces

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "Nil"
	Dogrula(sayi1)

	var sayi2 interface{} = 5
	Dogrula(sayi2)
}
