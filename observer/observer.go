package observer

import "fmt"

type SubjectInterface interface {
	Register(observer ObserverInterface)
	Remove(observer ObserverInterface)
	Notify(msg string)
}

type ObserverInterface interface {
	Update(msg string)
}

type Subject struct {
	observers []ObserverInterface
}

func (sub *Subject) Register(observer ObserverInterface) {
	sub.observers = append(sub.observers, observer)
}

func (sub *Subject) Remove(observer ObserverInterface) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, ob := range sub.observers {
		ob.Update(msg)
	}
}

type Observer1 struct{}

type Observer2 struct{}

func (o *Observer1) Update(msg string) {
	fmt.Println("Observer 1 : " + msg)
}

func (o *Observer2) Update(msg string) {
	fmt.Println("Observer 2 : " + msg)
}
