package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var result string

	if luckyNumber := 99; luckyNumber < 0 {

		result = "Number is less than zero."

	} else if luckyNumber == 0 {
		result = "Number is equal to zero."
	} else {
		result = "Number is greater than zero."
	}

	fmt.Println(result)

	// switch
	rand.Seed(time.Now().Unix())

	var switchResult string

	switch dow := rand.Intn(7) + 1; dow {

	case 1:
		switchResult = "It's Sunday"
		//fallthrough
	case 2:
		switchResult = "It's Monday"
	case 3:
		switchResult = "It's Tuesday"
	case 4:
		switchResult = "It's Wednesday"
	case 5:
		switchResult = "It's Thursday"
	case 6:
		switchResult = "It's Friday"
	case 7:
		switchResult = "It's Saturday"
	default:
		switchResult = "I am not sure about the day :("

	}

	fmt.Println(switchResult)

	// Loops / Iterations

	colors := []string{"Red", "Blue", "Black"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	} // for ends
	fmt.Println("*********************")
	for i := range colors {
		fmt.Println(colors[i])
	} // for ends

	fmt.Println("*********************")
	// ignore index

	for _, color := range colors {
		fmt.Println(color)
	} // for ends

	// instead of While loop, use for loop like this

	i := 0

	for i < len(colors) {
		fmt.Println(colors[i])

		i++
	} // for ends

	// c style goto for loop

	sum := 1

	for sum < 100 {
		sum += sum

		fmt.Println("Sum: ", sum)

		if sum >= 49 {
			goto theEnd
		}

	} // for ends

theEnd:

	fmt.Println("End of program.")
}
