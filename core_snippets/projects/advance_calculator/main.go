package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	number1 := getInput("Number 1:\n")
	number2 := getInput("Number 2:\n")

	var result float64

	switch operation := getOperation(); operation {

	case "+":
		result = addValues(number1, number2)
	case "-":
		result = subtractValues(number1, number2)
	case "*":
		result = multiplyValues(number1, number2)
	case "/":
		result = divideValues(number1, number2)
	default:
		panic("Invalid operation, please check your input numbers and try again.")
	} // switch ends

	// round the result for 2 digits in floating values
	result = math.Round(result*100) / 100

	fmt.Printf("This result is: %v\n\n\n", result)

} // main ends

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)

	input, _ := reader.ReadString('\n')

	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)

		panic(message)

	}

	return value

} // getInput ends

func getOperation() string {
	fmt.Println("Select mathematical operation (+, -, *, /) ")
	_operation, _ := reader.ReadString('\n')

	return strings.TrimSpace(_operation)
}

func addValues(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func subtractValues(num1 float64, num2 float64) float64 {
	return num1 - num2

}

func multiplyValues(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func divideValues(num1 float64, num2 float64) float64 {
	return num1 + num2
}
