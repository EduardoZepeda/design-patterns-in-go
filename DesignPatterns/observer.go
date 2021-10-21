package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		observers: make([]Observer, 0),
		name:      name,
		available: false,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Println("Item", i.name, "is available")
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, o := range i.observers {
		o.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(name string) {
	fmt.Println("Email to", e.id, "with item", name)
}

func main() {
	item := NewItem("Dog food")
	email1 := &EmailClient{id: "kyoko@kyoko.com"}
	item.register(email1)

	email2 := &EmailClient{id: "karenina@karenina.com"}
	item.register(email2)

	item.UpdateAvailable()
}
