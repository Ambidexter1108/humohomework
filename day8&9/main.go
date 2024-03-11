package main

import (
	"fmt"
	"log"
)

type user struct {
	Name string; 
	Surname string; 
	Age int; 
	Gender string;
	Email string; 
	Password string; 

}

func main() {
	log.Println()
 fmt.Println("\n-|-|-|-|-|-|-|-|-|-|-|-Программа запущена-|-|-|-|-|-|-|-|-|-|-|-")	
 fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~Добро пожаловать!~~~~~~~~~~~~~~~~~~~~~~~~")

	
 Klass := map[string]user{}

 Mentor := user{
				Name: "Shams",
				Surname: "Tabrezi",
				Age: 70,
				Gender: "male",
				Email: "shamsitabrezi1323@gmail.com",
				Password: "shamsuddin",

}
	// fmt.Printf("%+v\n", Mentor)

	Student := user{
				Name: "Jaloloddin",
				Surname: "Rumi",
				Age: 56,
				Gender: "male",
				Email: "jalol_rumi@gmail.com",
				Password: "balkhi",	
	}

	// fmt.Printf("%+v\n",Student)

	Klass["Jaloliddini Rumi"] = Student
	Klass["Shamsi Tabrezi"] = Mentor


// teacher1 := Klass["Boyazid"]
// fmt.Printf("%+v\n", teacher1)
// fmt.Printf("%+v\n", klass)
 fmt.Printf("*** В базе данных 2 человека:\n1}%+v\n2}%+v", Student.Name, Mentor.Name)

fmt.Println("\n*/*/* Выберите действие:")
fmt.Println("- 1 - Создание // Create")
fmt.Println("- 2 - Чтение // Read")
fmt.Println("- 3 - Модификация // Update")
fmt.Println(" -4 - Удаление // Delete")
	var action int
	fmt.Scanln(&action)
	switch action {
	case 1:
		Create(Klass)
	case 2:	
		Read(Klass)
	case 3:	
		Update(Klass, Student)
	case 4:	
		Delete(Klass)
	}
  

	fmt.Println("\n-|-|-|-|-|-|-|-|-|-|-|-Программа завершена-|-|-|-|-|-|-|-|-|-|-|-")	
	log.Println()
	fmt.Scan(&action)
}


func Create (Klass map[string]user) {
	_, ok := Klass["Boyazid"]
	if !ok {
		Klass["Boyazid"] = user {Name: "Boyazid", Surname: "Bastomi", Age: 90, Gender:"male", Email: "boyazid2211@gmail.com", Password: "guru"}
	}
	teacher1 := Klass["Boyazid"]
	log.Println("Создан новый персонаж:")
	fmt.Printf("%+v\n", teacher1)
	fmt.Printf("*** Теперь в базе данных 3 человека:\n%+v", Klass)
}
func Read (Klass map[string]user) {
	fmt.Println("*** В базе данных 2 человека:")
	fmt.Printf("%+v\n", Klass)
}
func Update (Klass map[string]user, Student user){
	Student2 := Student
 Student2.Age = 80
 Student2.Email = "jalol_balki@gmail.com"
 Student2.Password = "viceguru"
 fmt.Printf("Данные изменены успешно!\nДо:%+v\nПосле:%+v", Student, Student2)

} 

func Delete(Klass map[string]user) {
	delete(Klass, "Jaloliddini Rumi")
	fmt.Printf("Один персонаж удалён!\nТеперь в БД только один человек:%+v", Klass)
}

