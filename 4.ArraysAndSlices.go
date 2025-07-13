package main

import "fmt"

func main() {
	// Arrays......................................................................................................................
	var ages [3]int = [3]int{12, 23, 43}
	fmt.Println(ages, len(ages))

	names := [4]string{"John", "Kick", "Venom", "Michael"}
	fmt.Println(names, len(names))

	// Slicecs......................................................................................................................
	fmt.Println("\nAfter Slicing:")
	ages[1] = 100
	fmt.Println(ages, len(ages))

	names[1] = "Loona"
	fmt.Println(names, len(names))

	//Append Array....................................................................................................................
	var scores = []int{34, 35, 23, 12}
	scores[2] = 32
	scores = append(scores, 67)
	fmt.Println(scores, len(scores))

	// Slices By Range.................................................................................................................
	rangeOne := scores[1:3] // 1 Incusive, 3 Excusive (Like Py)
	rangeTwo := scores[2:]
	rangeThree := scores[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}
