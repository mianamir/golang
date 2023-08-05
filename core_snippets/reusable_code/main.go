package main

import "fmt"

func main() {
	greeting()

	fmt.Println(addTwoNumbers(23, 23))

	multiSum, multiCount := addAllValues(2, 3, 4, 5, 78, 89)

	fmt.Println(multiSum, multiCount)
}

func greeting() {
	fmt.Println("Welcome to Golang world.")
}

func addTwoNumbers(number1 int, number2 int) int {
	return number1 + number2
}

func addAllValues(values ...int) (int, int) {
	total := 0

	for _, v := range values {
		total += v
	}

	return total, len(values)
}
