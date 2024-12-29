package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Workshop1()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//loops.Demo4()
	//loops.Demo5()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//arrays.Demo4()
	//slices.Demo1()
	//slices.Demo2()
	//functions.SelamVer("Nil")
	//var sonuc = functions.Topla(2, 3)
	//fmt.Println(sonuc)

	/* sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)
	fmt.Println("Toplama: ", sonuc1)
	fmt.Println("Çıkarma: ", sonuc2)
	fmt.Println("Çarpma: ", sonuc3)
	//fmt.Printf("Bölme: %v", sonuc4) */

	/* sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	fmt.Println("Toplama: ", sonuc1)
	fmt.Println("Çıkarma: ", sonuc2)
	fmt.Println("Çarpma: ", sonuc3)
	fmt.Printf("Bölme: %v", sonuc4) */

	/* fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	fmt.Println(functions.ToplaVariadic(1, 4))
	fmt.Println(functions.ToplaVariadic())
	sayilar := []int{4, 6, 7, 0, 11}
	fmt.Println(functions.ToplaVariadic(sayilar...)) */

	//maps.Demo1()
	//ranges.Demo1()
	//ranges.Demo2()
	//ranges.Demo3()

	/* sayi := 20
	pointers.Demo1(&sayi)
	fmt.Println("Maindeki sayi: ", sayi) */

	/* sayilar := []int{1, 2, 3}
	pointers.Demo2(sayilar)
	fmt.Println("Maindeki sayı: ", sayilar[0]) */

	//structs.Demo1()
	//structs.Demo2()

	/* go goroutines.CiftSayilar()
	go goroutines.TekSayilar()
	//time.Sleep(time.Second)
	time.Sleep(5 * time.Second)
	fmt.Println("Main bitti") */

	/* ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)
	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Çarpım: ", carpim) */

	//interfaces.Demo1()
	//interfaces.Demo2()
	//defers.B()
	//defers.Test()
	//defers.Demo3()
	//error_handling.Demo1()
	//interfaces.Demo3()
	//error_handling.Demo2()
	//fmt.Println(error_handling.TahminEt2(102))
	//string_functions.Demo1()
	//string_functions.Demo2()
	//restful.Demo1()
	//restful.Demo2()
	//project.GetAllProducts()

	/* product, _ := project.AddProduct()
	fmt.Println(product)*/

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

	fmt.Println()
}
