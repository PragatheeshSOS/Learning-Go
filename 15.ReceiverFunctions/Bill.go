package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// New Bill.........................................................
func newBill(name string) Bill {
	new := Bill{
		name:  name,
		items: map[string]float64{"Apple": 65.54, "Banana": 78.5, "Orange": 49.02},
		tips:  1.94,
	}
	return new
}

// Formatting The Bill..............................................
func (formatedBill Bill) format() string {
	res := "Bill Details\n"
	var total float64 = 0
	for item, price := range formatedBill.items {
		res += fmt.Sprintf("%-25v : %0.2f\n", item, price)
		total += price
	}
	// Tips.............................................
	res += fmt.Sprintf("%-25v : %0.2f\n", "Tips", formatedBill.tips)
	total += formatedBill.tips

	//Total.............................................
	res += fmt.Sprintf("%-25v : %0.2f\n", "Total", total)
	return res
}
