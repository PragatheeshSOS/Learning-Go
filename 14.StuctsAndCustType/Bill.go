package main

type Bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// New Bill.........................................................
func newBill(name string) Bill {
	new := Bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
	}
	return new
}
