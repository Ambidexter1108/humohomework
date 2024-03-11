package main

import (
	"fmt"
	"log"
	"unicode/funcs"
)

func main() {
	log.Println()
	fmt.Println("*|*|*|*|*|*|*|*|*|*|Программа запущена|*|*|*|*|*|*|*|*|*|*")

	fmt.Println("Выберите задание:")

	fmt.Println("- Задание 1: Конвертация строк и байтов")
	fmt.Println("- Задание 2: Шифр Цезаря // Cдвиг на 3 буквы")
	fmt.Println("- Задание 3: Исправить и заменить '+' на 'пробел'")

	var task uint8
	fmt.Scanln(&task)

	switch task {
	case 1: 
	
	fmt.Println("Выберите опцию:")
	fmt.Println("1) Конвертация строк в байт")
	fmt.Println("2) Конвертация байт в строк")

	var task1 uint8
	fmt.Scanln(&task1)

			if task1 == 1 {
				funcs.StrToUint()
			} else if task1 == 2 {
				funcs.UintToStr()
			}


	case 2:
		funcs.Caesar()

	case 3:
		funcs.Antiplus()
}





	fmt.Println("*|*|*|*|*|*|*|*|*|*|Программа завершена|*|*|*|*|*|*|*|*|*|*")
	log.Println()
fmt.Scanln(&task)
}
