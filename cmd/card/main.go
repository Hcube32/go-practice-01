package main

import "fmt"

func main() {
	const name_course = "Бэкенд-разработка на языке Go"
	var name string = "Семён"
	var group string = "ЭПБО-01-2026"
	count := 1
	enjoyment := true

	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", name_course)
	fmt.Printf("Завершено работ: %d\n", count)
	fmt.Printf("Интерес к бэкенду: %t", enjoyment)

}
