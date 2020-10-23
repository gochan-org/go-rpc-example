package rpccommon

import "fmt"

// Main does most of the changes
type Main int

func (m *Main) HelloWorld(a *int, b *int) error {
	fmt.Println("Hello, RPC!")
	return nil
}

// ChangeColor changes the RGB values of c1 to c2
func (m *Main) ChangeColor(colors []Color, reply *int) error {
	fmt.Printf("colors: %#v\n", colors)
	// c1.Red = c2.Red
	// c1.Green = c2.Green
	// c1.Blue = c2.Blue
	return nil
}

// Color represents a generic RGB color object
type Color struct {
	Red   int
	Green int
	Blue  int
}
