package main

import "fmt"

func main() {
	a := 20
	b := 25
	c := 35

	sum := a + b + c
	avg := (float64(a) + float64(b) + float64(c)) / 3

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее: %.2f\n", avg)
}
