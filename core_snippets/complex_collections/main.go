package main

import (
	"fmt"
	"sort"
)

func main() {

	/**

	1- Collection Types
	2- Memory management in Golang
	3- Pointers
	4- Arrays

	**/

	// Pointers

	fmt.Println("** Pointers **")
	number1 := 23

	var ptr = &number1

	fmt.Println("Value of ptr is: ", ptr)

	number2 := 234

	pointer2 := &number2

	fmt.Println("Value of ptr2 is: ", pointer2)

	*pointer2 = *pointer2 / 23

	fmt.Println("Value of pointer2 is: ", *pointer2)
	fmt.Println("Value of number2 is: ", number2)

	// Arrays
	fmt.Println("** Arrays **")
	var colors [3]string

	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println("Total colors array length: ", len(colors))

	var numbers = [3]float32{23.4, 45.6, 67.8}

	fmt.Println(numbers)
	fmt.Println("Total numbers array length: ", len(numbers))

	// Slices
	fmt.Println("** Slices **")

	var newColorsArray = []string{"Orange", "Black", "Pink", "Yellow", "Green", "Gray"}

	fmt.Println(newColorsArray)

	// add item in the Slice
	newColorsArray = append(newColorsArray, "Brown")
	fmt.Println(newColorsArray)

	// Remove first item from the Slice
	newColorsArray = append(newColorsArray[1:len(newColorsArray)])
	fmt.Println(newColorsArray)

	// Remove last item from the Slice
	newColorsArray = append(newColorsArray[1 : len(newColorsArray)-1])
	fmt.Println(newColorsArray)

	// declare the Slice with make()

	numbersArray := make([]int, 3, 3)

	numbersArray[0] = 312
	numbersArray[1] = 29
	numbersArray[2] = 25

	fmt.Println(numbersArray)

	numbersArray = append(numbersArray, 45)

	fmt.Println(numbersArray)

	sort.Ints(numbersArray)
	fmt.Println(numbersArray)

	// Maps
	fmt.Println("** Maps **")

	cities := make(map[string]string)

	cities["Ber"] = "Berlin"
	cities["Mun"] = "Munich"
	cities["Erf"] = "Erfurt"
	cities["Col"] = "Colone"

	fmt.Println(cities)

	munich := cities["Mun"]

	fmt.Println(munich)

	// remove item from map
	delete(cities, "Col")
	fmt.Println(cities)

	// loop on map

	for k, v := range cities {

		fmt.Printf("%v : %v\n", k, v)
	} // for ends

	// get keys from map and sort them

	mapKeys := make([]string, len(cities))

	i := 0

	for k := range cities {
		mapKeys[i] = k

		i++
	} // for ends

	fmt.Println(mapKeys)
	sort.Strings(mapKeys)
	fmt.Println(mapKeys)

	// use iterative loop

	for i := range mapKeys {
		fmt.Println(cities[mapKeys[i]])
	} // for ends

	// Structs
	fmt.Println("** Structs **")

}
