package main

import "fmt"

func main() {

	// TODO: print values in same or new lines
	fmt.Print("Testing print func without new line.")

	fmt.Println("Testing print func with new line.")

	// TODO: print the slice data
	items := []int{12, 13, 14, 15, 16}

	length, err := fmt.Println(items)

	fmt.Println("Slice array number of bytes", length, err)
}
