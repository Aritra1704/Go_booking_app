package main

import "fmt" // format package

// REF: https://www.youtube.com/watch?v=yyUHQIec83I

func main() {
	var conferenceName = "Go Conference"
	const confernceTickets = 50
	var remainingTickets = confernceTickets

	fmt.Printf("Welcome to %v our booking application\n", conferenceName)
	fmt.Println("We have a total of", confernceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here")

}
