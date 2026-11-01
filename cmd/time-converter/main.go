package main

import "fmt"

func main() {
	var digit int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&digit)
	if err != nil {
		fmt.Println("Только числа")
		return
	}
	if digit < 0 {
		fmt.Println("Время не может быть отрицательнным")
	} else {
		hours := digit / 60
		mins := digit % 60

		fmt.Printf("Полных часов: %d\n", hours)
		fmt.Printf("Полных минут: %d\n", mins)

	}

}
