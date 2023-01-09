package main

import "fmt" // format package

// REF: https://www.youtube.com/watch?v=yyUHQIec83I
// 53:59
func main() {
	var conferenceName string = "Go Conference"
	const confernceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v our booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confernceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The first value %v\n", bookings[0])
	fmt.Printf("The array type %T\n", bookings)
	fmt.Printf("The array length %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaning for %v\n", remainingTickets, conferenceName)
}
