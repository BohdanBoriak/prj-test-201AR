package main

import (
	"fmt"
	"math/rand"
	"prj-test/domain"
	"strconv"
	"time"
)

const (
	totalPoints       = 100
	pointsPerQuestion = 20
)

func main() {
	fmt.Println("Вітаємо у грі MathCore!")

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			play()
		case "2":
			fmt.Println("Рейтинг в розробці -_-")
		case "3":
			return
		default:
		}
	}
}

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func play() {
	timeStart := time.Now()
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У Вас %v очок!\n", myPoints)
			} else {
				fmt.Println("НЕ ПРАВИЛЬНО!")
			}
		}
	}
	timeFinish := time.Now()
	timeSpent := timeFinish.Sub(timeStart)

	fmt.Printf("Ваш час: %v\n", timeSpent)
	fmt.Print("Введіть своє ім'я: ")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Name:      name,
		TimeSpent: timeSpent,
	}
}
