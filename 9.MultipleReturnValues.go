package main

// import (
// 	"fmt"
// 	"strings"
// )

// func getStartingLetter(start string) (string, string) {
// 	fullName := strings.ToUpper(start)
// 	names := strings.Split(fullName, " ")
// 	var res []string
// 	for _, value := range names {
// 		res = append(res, value[:1])
// 	}
// 	if len(names) > 1 {
// 		return res[0], res[1]
// 	}
// 	return res[0], "_"
// }

// func main() {
// 	firstName1, secondName1 := getStartingLetter("wonda vision")
// 	fmt.Printf("First Name Letter Is %v And Second Name Letter Is %v\n", firstName1, secondName1)

// 	firstName2, secondName2 := getStartingLetter("tony stark")
// 	fmt.Printf("First Name Letter Is %v And Second Name Letter Is %v\n", firstName2, secondName2)

// 	firstName3, secondName3 := getStartingLetter("thor")
// 	fmt.Printf("First Name Letter Is %v And Second Name Letter Is %v\n", firstName3, secondName3)
// }
