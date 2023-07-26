package main

import "fmt"

func main() {

	car := Car{"2023 BMW iX", 14400.56}

	fmt.Println(car)
	fmt.Printf("%+v\n", car)
	fmt.Printf("Modal %v\nPrice %v\n", car.Modal, car.Price)

}

type Car struct {
	Modal string
	Price float64
}
