package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	// TODO: обработайте new, in_progress, done через switch.
	// TODO: предусмотрите ветку default для незнакомого статуса.

	switch status {
	case "new":
		return "Новая задача"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Незнакомый статус"
	}
	//return "Добавьте перевод статуса"
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	count := 0
	for _, stat := range statuses {
		if stat == target {
			count++
		}
	}
	return count
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done", "blocked"}
	// TODO: добавьте ещё один статус new через append.
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, stat := range statuses {
		counts[stat]++
	}

	fmt.Println("Новых:", counts["new"])
	fmt.Println("В работе:", counts["in_progress"])
	fmt.Println("Завершено:", counts["done"])

	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.
}
