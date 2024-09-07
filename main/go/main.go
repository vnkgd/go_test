package main

import "fmt"

// func main() {
// 	var myName = "Jason"
// 	fmt.Println("My name is", myName)
// 	var name string = "John"
// 	fmt.Println("name =", name)
// 	userName := "admin"
// 	fmt.Println("username =", userName)
// 	var sum int
// 	fmt.Println("sum is =", sum)
// 	a, b := 5, 7
// 	fmt.Println("число 1 =", a, "число 2 =", b)
// 	a, h := 9, 4
// 	fmt.Println("число 4 =", a, "число 4 =", h)
// 	sum2 := a + b + h
// 	fmt.Println("sum2 = ", sum2)
// }

func main() {
	//одна переменная
	var favCol = "Фиолетовый"
	fmt.Println("Цвет = ", favCol)

	//две переменные
	bithDate, ageInYears := 2024, 124
	fmt.Println("Год рождения = ", bithDate, "Возраст = ", ageInYears)

	//блочное присвоение
	var (
		firstInitial  = "A"
		sunameInitial = "П"
	)
	fmt.Println("Инициалы = ", firstInitial, sunameInitial)

	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("Возраст в днях =", ageInDays)
}
