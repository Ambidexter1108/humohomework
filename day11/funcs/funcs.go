package funcs

import (
	"fmt"
	"strings"
)

func StrToUint() {
	fmt.Println("Введите строку:")

	var text1 string
	fmt.Scanln(&text1)

	num := []byte(text1)
	fmt.Println(text1, "=", num)

}

func UintToStr() {
	fmt.Println("Введите байт:")
	var num2 []byte

	text2 := string(num2)
	fmt.Println(num2, "=", text2)
}

func Caesar() {
	fmt.Println("Введите слово in English:")

	var text string
	fmt.Scanln(&text)

	num := []byte(text)
	fmt.Println(num)

	for _, c := range num {

		if c > 119 || (c >= 88 && c <= 90) {
			fmt.Printf("%c = %c\n", c, c-23)
		} else if c < 119 {
			fmt.Printf("%c = %c\n", c, c+3)
		}
	}
}

func Antiplus() {
	fmt.Println("Введите+текст+таким+образом:")

	var text3 string
	fmt.Scanln(&text3)

	noPlus := strings.Replace(text3, "+", " ", -1)
	fmt.Println(noPlus)
}
