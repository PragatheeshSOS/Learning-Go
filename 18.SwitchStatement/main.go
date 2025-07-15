package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, read *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := read.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(new Bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose Option (A - Add Item, S - Save Bill, T - Add Tip): ", reader)
	switch option {
	case "A":
		name, _ := getInput("Enter Item Name: ", reader)
		price, _ := getInput("Enter Item Price: ", reader)
		fmt.Println(name, price)
	case "S":
		tip, _ := getInput("Enter Tips: ", reader)
		fmt.Println(tip)
	case "T":
		fmt.Println("You Choose T.")
	default:
		fmt.Println("Invalid Option. Try Again..!")
		promptOptions(new)
	}
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create A New Bill Name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create A New Bill Name: ", reader)
	bill := newBill(name)
	fmt.Println("Created Bill Name:", bill.name)
	return bill
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
