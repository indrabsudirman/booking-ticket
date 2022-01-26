package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("ConferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v Booking ticket here\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get all your ticket here")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their  name
		fmt.Println("Please enter your fisrt name : ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address : ")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets : ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}

}
