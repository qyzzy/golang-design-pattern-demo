package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	o1, o2 := &Observer1{}, &Observer2{}
	sub.Register(o1)
	sub.Register(o2)
	sub.Notify("msg")
	o1.Update("1")
	o2.Update("2")
}
