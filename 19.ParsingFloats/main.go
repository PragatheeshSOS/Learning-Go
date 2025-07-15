package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		parsingPrice, price_err := strconv.ParseFloat(price, 64)
		if price_err != nil {
			fmt.Println("The Price Must Be A Number.")
			promptOptions(new)
		}
		new.addItems(name, parsingPrice)
		fmt.Println("Item Added:", name, price)
		promptOptions(new)
	case "T":
		tip, _ := getInput("Enter Tips: ", reader)
		parsingTip, tip_err := strconv.ParseFloat(tip, 64)
		if tip_err != nil {
			fmt.Println("The Price Must Be A Number.")
			promptOptions(new)
		}
		new.updateTips(parsingTip)
		fmt.Println("Tips Updated:", tip)
		promptOptions(new)
	case "S":
		fmt.Println("Bill Saved:", new)
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
