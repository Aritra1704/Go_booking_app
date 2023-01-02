package main

import "fmt" // format package

// REF: https://www.youtube.com/watch?v=yyUHQIec83I

func main() {
	var conferenceName = "Go Conference"
	const confernceTickets = 50
	var remainingTickets = confernceTickets

	fmt.Printf("Welcome to %v our booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confernceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}
