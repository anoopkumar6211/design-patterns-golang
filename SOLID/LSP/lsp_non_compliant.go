/*
The Square is a subtype of Rectangle, but the logic for Area is incorrect when treating a square as a rectangle. The Square class does not behave like a Rectangle as expected, violating LSP.
The user expects to be able to use Square wherever Rectangle is used, but this doesn't work as expected.

*/

package main

import "fmt"

// Rectangle is the base struct.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square inherits from Rectangle but incorrectly modifies its behavior.
type Square struct {
	Rectangle
}

// SetSide sets the width and height of the square to be the same.
func (s *Square) SetSide(side float64) {
	s.Width = side
	s.Height = side
}

func main() {
	var rect Rectangle = Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", rect.Area()) // Output: Rectangle Area: 50

	// Substituting Square for Rectangle violates LSP as the behavior is unexpected.
	var square Rectangle = Square{}
	square.Width = 5 // Trying to set the width only for a square is incorrect
	fmt.Println("Square Area (incorrect):", square.Area()) // Output: Square Area: 0
}
