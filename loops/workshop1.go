package loops

import "fmt"

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		} else if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz!")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""

	if tahminEdilenSayi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper!"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena Değil"
	}

	fmt.Println("Tabrikler! Bildiniz.")
	fmt.Printf("Tahmin sayısı: %v, %v", tahminSayisi, basariDurumu)

}
