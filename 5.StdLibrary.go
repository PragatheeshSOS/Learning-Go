package main

// import (
// 	"fmt"     // Printing & Formatting Package
// 	"sort"    // Sort Package
// 	"strings" //String Package
// )

// func main() {
// 	// Strings Package.......................................................................................
// 	fmt.Println("String Packages")
// 	greetings := "Hello World..!"
// 	fmt.Println(strings.Contains(greetings, "Hello")) //Output = True/False. Verifies Given Value Is Present Or Not In The String.
// 	fmt.Println(strings.ReplaceAll(greetings, "Hello ", "Hello To "))
// 	fmt.Println(strings.ToUpper(greetings))
// 	fmt.Println("Original String Is", greetings)
// 	fmt.Println(strings.Index(greetings, "W"))
// 	fmt.Println(strings.Split(greetings, " "))

// 	// Sort Package.......................................................................................
// 	fmt.Println("Sort Packages")
// 	ages := []int{54, 34, 76, 86, 12, 13, 11}
// 	sort.Ints(ages)
// 	fmt.Println(ages)

// 	index := sort.SearchInts(ages, 76)
// 	fmt.Println("Index Of Value After Sorting:", index) // If Element Not In List. Then, It Returns Len(list)+1.

// 	names := []string{"John", "Kick", "Venom", "Michael"}
// 	sort.Strings(names)
// 	fmt.Println(names)

// 	fmt.Println(sort.SearchStrings(names, "Kick"))
// }
