package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	conferenceTickets := 50
	remainingTickets := 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Ask user for their first name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	// Ask user for their last name
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	// Ask user for their email address
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	// Ask user for the number of tickets
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will received a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
