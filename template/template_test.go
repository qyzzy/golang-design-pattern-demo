package template

import "testing"

func TestA_Play(t *testing.T) {
	a := &A{}
	a.InitGame()
	a.Play()
	a.Stop()
}
