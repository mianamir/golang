package main

import "fmt"

// Go compile errors to enforce better code quality
// In this way we can avoid possible dead code, code that is never used in the code
// Static type of Go has these features
// More robust, reduces the likelihood of errors
// Help developers to catch type mismatches sooner (at compile time)
// Go discover mistakes at compile time Not at runtime


func main(){

	// type inference in Go

	conferenceName := "Go Conference" // var conferenceName string = "Go Conference"

	const conferenceTickets int = 50 // constant value

	var remainingTickets uint = 50 // automatic error handling for -ve values

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", 
								conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v, are still available.\n", 
								conferenceTickets, 
								remainingTickets)

	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) // add variable pointer

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName) // add variable pointer


	fmt.Println("Enter your email:")
	fmt.Scan(&email) // add variable pointer


	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets) // add variable pointer


	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will a receive confirmation email at %v\n",
							firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)











	

	


}