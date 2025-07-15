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
	fmt.Println(option)
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
