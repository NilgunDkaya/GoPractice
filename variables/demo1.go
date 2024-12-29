package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + 100*kdv/100)

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25.2
	fmt.Println(kdv3)
	fmt.Printf("Veri Tipi: %T", kdv3)
	fmt.Println()

	//var durum bool = false

	var metin1 string = "Nil"
	var metin2 string = "Demir"

	var durum bool = (metin1 != metin2)
	fmt.Println(durum)
	fmt.Println(!durum)
}
