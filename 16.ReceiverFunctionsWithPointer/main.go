package main

import "fmt"

func main() {
	myBill := newBill("Mona's Bill")
	fmt.Println(myBill)
	fmt.Println(myBill.format())

	myBill.addItems("Tomato", 16.5)
	myBill.addItems("Rice", 106.45)
	myBill.addItems("Wonder Cake", 74.60)
	myBill.addItems("7Up 1/2Ltr", 34.98)
	myBill.updateTips(10.37)
	fmt.Println(myBill.format())
}
