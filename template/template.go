package template

import "fmt"

type Game interface {
	InitGame()
	Play()
	Stop()
}

type A struct{}

func (a *A) InitGame() {
	fmt.Println("A init")
}

func (a *A) Play() {
	fmt.Println("A play")
}

func (a *A) Stop() {
	fmt.Println("A stop")
}

type B struct{}

func (b *B) InitGame() {
	fmt.Println("B init")
}

func (b *B) Play() {
	fmt.Println("B play")
}

func (b *B) Stop() {
	fmt.Println("B stop")
}
