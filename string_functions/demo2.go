package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "NilNil"
	fmt.Println(s.HasPrefix(isim, "Nil"))
	fmt.Println(s.HasSuffix(isim, "il"))
	fmt.Println(s.HasSuffix(isim, "nil"))
	fmt.Println(s.Index(isim, "il"))
	harfler := []string{"n", "i", "l", "n", "i", "l"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "-", 1))
	fmt.Println(s.Replace(sonuc, "*", "-", 3))
	fmt.Println(s.Replace(sonuc, "*", "-", -1))

	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 3))
}
