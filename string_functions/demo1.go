package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "NilNil"
	fmt.Println(s.Count(isim, "i"))
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "i"))
	fmt.Println(s.Index(isim, "N"))
	fmt.Println(s.Index(isim, "a"))

	sonuc := s.Index(isim, "a")
	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
