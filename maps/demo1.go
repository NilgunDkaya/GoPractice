package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["Table"] = "Masa"
	sozluk["Book"] = "Kitap"
	sozluk["Pencil"] = "Kalem"

	fmt.Println(sozluk["Book"])
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	delete(sozluk, "Book")
	fmt.Println("Eleman sayısı: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["Table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)

	deger2, varMi := sozluk["Lamb"]
	fmt.Println(deger2)
	fmt.Println("Listede olma durumu: ", varMi)

	sozluk2 := map[string]string{"Glass": "Bardak",
		"Microphone": "Mikrofon"}

	fmt.Println(sozluk2)
}
