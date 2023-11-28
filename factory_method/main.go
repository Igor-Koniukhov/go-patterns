package main

import "fmt"

// Product define the product interface
type Product interface {
	Use() string
}

// ConcreteProductA и ConcreteProductB — concrete implementation Product
type ConcreteProductA struct{}
type ConcreteProductB struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A"
}

func (p *ConcreteProductB) Use() string {
	return "Product B"
}

// Creator define abstract factory method
type Creator interface {
	CreateProduct(t string) Product
}

// ConcreteCreator implements Creator
type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct(t string) Product {
	switch t {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	creator := &ConcreteCreator{}
	productA := creator.CreateProduct("A")
	fmt.Println(productA.Use())

	productB := creator.CreateProduct("B")
	fmt.Println(productB.Use())
}
