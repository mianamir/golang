package main

import "fmt"

func main() {
	poddle := Dog{"Poddle", 10, "Woof"}

	fmt.Println(poddle)

	fmt.Printf("%+v\n", poddle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poddle.Breed, poddle.Weight)

	poddle.Speak()

	poddle.Sound = "Arf"

	poddle.Speak()

	poddle.SpeakThreeTimes()
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how the dog speaks 3 times
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)

	fmt.Println(d.Sound)
}
