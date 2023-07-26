package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Value 1: ")

	input1, _ := reader.ReadString('\n')

	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number.")
	}

	fmt.Println("Value 2: ")

	input2, _ := reader.ReadString('\n')

	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number.")
	}

	twoNumbersSum := float1 + float2

	sumResult := math.Round(twoNumbersSum*100) / 100

	fmt.Printf("The sum of %v and %v is %v \n\n", float1, float2, sumResult)

	fmt.Println("**************")

}
