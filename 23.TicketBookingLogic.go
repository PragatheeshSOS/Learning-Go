package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var name string // "Mona"
// 	var age int16   // 18
// 	var mail string
// 	var totalTickets int16 = 100
// 	var reserveTickets int16
// 	for {
// 		fmt.Print("Enter Name: ")
// 		fmt.Scan(&name)
// 		fmt.Print("Enter Age: ")
// 		fmt.Scan(&age)
// 		fmt.Print("Enter Email: ")
// 		fmt.Scan(&mail)
// 		fmt.Print("Enter Number Tickets To Reserve: ")
// 		fmt.Scan(&reserveTickets)
// 		fmt.Println("Loading....!")

// 		time.Sleep(5 * time.Second) // Slows Down The Response.......
// 		var remainingTickets int16 = totalTickets - reserveTickets
// 		fmt.Println("\n\nBooking Successfull..!")
// 		fmt.Printf("Name : %v\n", name)
// 		fmt.Printf("Age : %d\n", age)
// 		fmt.Printf("Email : %v\n", mail)
// 		fmt.Printf("Number Of Booked Tickets : %v\n", reserveTickets)
// 		fmt.Println("Available Tickets: ", remainingTickets)

// 		if remainingTickets <= 0 {
// 			print("Booking Closed...!")
// 			break
// 		}
// 	}
// }
