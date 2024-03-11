package main

import (
	"fmt"
	"log"
	"math/rand"
)

type user struct {
	Name         string
	Surname      string
	Family_Name  string
	Age          uint8
	Phone_number int
	Status       string
	Card_number  int
	Pin          int
	Money        int
}

// интерфейс с методами
type singleQr interface {
	setName(string, string, string)
	setAge(uint8)
	setPhone_number(int)
	setStatus(string)
	setMoney(int)
	setCard_number(int)
	setPin(int)
	changePin(int)
	check_balance(int) int
}

func addCard(DC singleQr) {
	DC.setCard_number(9762000022110011)
	DC.setPin(2211)
	DC.setMoney(999)
}

// выход
// func Quit(Q singleQr) {
// 	Q.setStatus("Не идентифицирован")
// }

// идентификация
func iDentificate(C singleQr) {
	C.setStatus("Идентифицирован")
}

// платеж
func Payment(money int) (bonus int) {
	return
}

// // проверка баланса
// func cheQue(Ch singleQr) {
// 	Ch.check_balance()
// }

func main() {

	log.Println()
	fmt.Println("^^^^^^^^^^^^Программа запущена^^^^^^^^^^^^")
	fmt.Println("Добро пожаловать!")

	fmt.Println("Выберите операцию:")
	fmt.Println("- 1. Добавить пользователя / Регистрация")

	var entry int
	fmt.Scanln(&entry)
	switch entry {
	case 1:
		fmt.Println("Введите имя:")
		var n string
		fmt.Scanln(&n)
		var newbie *user
		newbie = new(user)
		newbie.Name = n
		newbie.Status = "Не идентифицирован"
		log.Printf("Вход выполнен:%q, %q", n, newbie.Status)
	}
	fmt.Println("<<< Чтобы пользоваться услугами электронного кошелька, нужно идентифицироваться!")
	fmt.Println("Выберите действие:")
	fmt.Println("- 1.Идентификация")

	var ident int
	fmt.Scan(&ident)
	var Data *user
	switch ident {
	case 1:
		var newUser *user
		newUser = new(user)

		fmt.Println("* * * Введите имя, фамилию и отчество:")
		var n, s, f string
		fmt.Scan(&n, &s, &f)
		fmt.Scanln(&n, &s, &f)

		newUser.Name = n
		newUser.Surname = s
		newUser.Family_Name = f

		fmt.Println("* * * Введите ваш возраст:")
		var age uint8
		fmt.Scanln(&age)
		newUser.Age = age

		fmt.Println("* * * Введите номер вашего телефона:")
		var ph int
		fmt.Scanln(&ph)
		newUser.Phone_number = ph

		iDentificate(newUser)

		log.Printf("Идентификация прошла успешно!:% +v\n", newUser)
		Data = newUser
	}

	fmt.Println("* * * Нажмите 1 и enter чтобы добавить карту")
	var dcard int
	fmt.Scanln(&dcard)
	switch dcard {
	case 1:
		addCard(Data)
		log.Println("Карта добавлено успешно!")
	}

	fmt.Println("* * * Выберите опцию:\n",
		"- 1. Оплатить электричество\n",
		"- 2. Перевести деньги\n",
		"- 3. Проверить баланс",
	)
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		var elec int
		elec = rand.Intn(400)
		fmt.Println("* * * Электричество заработало", elec , "TJS")
		fmt.Println("Сколько вы хотите заплатить?")
		var payment int
		fmt.Scanln(&payment)
		Data.bill(Data.Money, payment)
		log.Println("Вы заплатили за электричество", payment, "TJS.", "Ваш баланс(При проверке баланса взимается небольшая сумма):", Data.check_balance(Data.Money), "TJS")
	case 2:
		fmt.Println("* * * Введите имя получателя")
		var rname string
		fmt.Scanln(&rname)
		fmt.Println("Введите номер получателя")
		var receiver int
		fmt.Scanln(&receiver)
		fmt.Println("Введите сумму")
		var sum int
		fmt.Scanln(&sum)
		Data.transaction(Data.Money, sum)
		log.Println("Вы перевели", sum, "TJS.", "Получатель:", rname, ".", "Ваш баланс(При проверке баланса взимается небольшая сумма):", Data.check_balance(Data.Money), "TJS")
	case 3:
		log.Println("Ваш баланс(При проверке баланса взимается небольшая сумма):", Data.check_balance(Data.Money), "TJS")
	}
	
	 fmt.Println("^^^^^^^^^^^^Программа завершена^^^^^^^^^^^^")
	 log.Println()
	 var quit string
	fmt.Scanln(&quit)
}

func (businessman *user) transaction(money, receiver int) (after int) {
	businessman.Money = money
	after = money - receiver
	businessman.Money = after
	return businessman.Money
}

func (businessman *user) bill(money, receiver int) (after int) {
	businessman.Money = money
	after = money - receiver
	businessman.Money = after
	return businessman.Money
}
func (businessman *user) check_balance(money int) (balance int) {
	businessman.Money = money
	balance = money - rand.Intn(3)
	businessman.Money = int(balance)
	return businessman.Money
}

func (businessman *user) setName(name, surname, familyname string) {
	businessman.Name = name
	businessman.Surname = surname
	businessman.Family_Name = familyname
}

func (aligarh *user) setCard_number(card_number int) {
	aligarh.Card_number = card_number

}

// changePin implements singleQr.
func (vip *user) changePin(pin int) {
	vip.Pin = pin
}

// setAge implements singleQr.
func (enterpreneur *user) setAge(age uint8) {
	enterpreneur.Age = age
}

// setMoney implements singleQr.
func (vip *user) setMoney(money int) {
	vip.Money = money
}

// setPhone_number implements singleQr.
func (enterpreneur *user) setPhone_number(phone_number int) {
	enterpreneur.Phone_number = phone_number
}

// setPin implements singleQr.
func (aligarh *user) setPin(pin int) {
	aligarh.Pin = pin
}

// setStatus implements singleQr.
func (newuser *user) setStatus(status string) {
	newuser.Status = status
}
