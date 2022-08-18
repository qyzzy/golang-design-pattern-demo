package factory

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length, width float64
}

func (s *Rectangle) Area() float64 {
	return s.length * s.width
}

type Circle struct {
	radius float64
}

func (s *Circle) Area() float64 {
	return (s.radius * s.radius) * math.Pi
}

func (s *Circle) SetRadius(radius float64) {
	s.radius = radius
}

type Square struct {
	length float64
}

func (s *Square) Area() float64 {
	return s.length * s.length
}

type Factory struct {
}

type shapeType string

func (factory *Factory) GetShape(t shapeType) Shape {
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
