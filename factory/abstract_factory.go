package factory

import "fmt"

type colorType string

type AbstractFactory interface {
	GetShape(t shapeType) Shape
	GetColor(t colorType) Color
}

type Color interface {
	Fill()
}

type Blue struct {
}

func (c *Blue) Fill() {
	fmt.Println("blue fill")
}

type Green struct {
}

func (c *Green) Fill() {
	fmt.Println("green fill")
}

type AbsFactory struct {
}

func (factory *AbsFactory) GetShape(t shapeType) Shape {
	if t == "" {
		return nil
	}
	switch t {
	case "circle":
		return new(Circle)
	case "rectangle":
		return new(Rectangle)
	case "square":
		return new(Square)
	}
	return nil
}

func (factory *AbsFactory) GetColor(t colorType) Color {
	if t == "" {
		return nil
	}
	switch t {
	case "blue":
		return new(Blue)
	case "green":
		return new(Green)
	}
	return nil
}
