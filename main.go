package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("Rectangle Width:", rect.Width)
	fmt.Println("Rectangle Height:", rect.Height)
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}
