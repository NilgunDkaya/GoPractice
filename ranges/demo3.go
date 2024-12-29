package ranges

import "fmt"

func Demo3() {
	sozluk := map[string]string{"Book": "Kitap", "Table": "Masa"}
	for k, v := range sozluk {
		fmt.Print(k)
		fmt.Println(" ", v)
	}
}
