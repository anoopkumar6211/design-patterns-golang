/*
Interface: The Shape interface defines a contract for all shapes, with an Area method. Both Rectangle and Square implement this interface, providing their own implementations of the Area method.
LSP-Compliant: Now, a Square can be substituted for a Rectangle as they both conform to the Shape interface. The behavior remains correct when using Square or Rectangle through the Shape abstraction.

*/

package main

import "fmt"

// Shape interface defines an abstraction for different shapes.
type Shape interface {
	Area() float64
}

// Rectangle is a shape that implements the Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square is a different shape that also implements the Shape interface.
type Square struct {
	Side float64
}

// Area returns the area of a square.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	square := Square{Side: 4}

	printArea(rectangle) // Output: Area: 50.00
	printArea(square)    // Output: Area: 16.00
}
