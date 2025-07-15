package main

// import (
// 	"fmt"
// 	"math"
// )

// // Interface............................................................
// type shape interface {
// 	area() float64
// 	circumference() float64
// }

// type square struct {
// 	length float64
// }

// type circle struct {
// 	radius float64
// }

// // Square Method.........................................................
// func (Shape square) area() float64 {
// 	return Shape.length * Shape.length
// }

// func (Shape square) circumference() float64 {
// 	return Shape.length * 4
// }

// // Circle Method...........................................................
// func (Shape circle) area() float64 {
// 	return math.Pi * Shape.radius * Shape.radius
// }

// func (Shape circle) circumference() float64 {
// 	return 2 * math.Pi * Shape.radius
// }

// func Printing(Shape shape) {
// 	fmt.Printf("Area Of %T Is: %0.2f\n", Shape, Shape.area())
// 	fmt.Printf("Circumferance Of %T Is: %0.2f\n", Shape, Shape.circumference())
// 	fmt.Println("-------------------------------")
// }

// func main() {
// 	Shapes := []shape{
// 		square{length: 15.2},
// 		circle{radius: 7.5},
// 		circle{radius: 12.3},
// 		square{length: 4.9},
// 	}
// 	fmt.Println("\n------------------------------")
// 	for _, value := range Shapes {
// 		Printing(value)
// 	}
// }
