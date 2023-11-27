package main

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() string {
	return "Hello"
}

type Decorator struct {
	Component
}

type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) Operation() string {
	return d.Component.Operation() + ", World!"

}

type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) Operation() string {
	return d.Component.Operation() + " Nice to see you!"
}

func main() {
	var comp Component = &ConcreteComponent{}

	comp = &ConcreteDecoratorA{Decorator{comp}}
	comp = &ConcreteDecoratorB{Decorator{comp}}

	fmt.Println(comp.Operation())

}
