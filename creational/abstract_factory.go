package creational

// Abstract Factory is a creational design pattern that lets you produce families of related objects without
// specifying their concrete classes.

// In this example we have an abstract factory interface SportsFactory, then we define a function getSportsFactory that
// returns the concrete factory of the given family of related objects.

import "fmt"

// SportsFactory is the interface of the Abstract Factory.
type SportsFactory interface {
	makeShoe() Shoe
	makeShirt() Clothing
}

func getSportsFactory(brand string) (SportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidasFactory{}, nil
	case "nike":
		return &nikeFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown brand %s", brand)
	}
}

// adidasFactory is the Concrete Factory, it implements the abstract factory interface.
type adidasFactory struct{}

func (a *adidasFactory) makeShoe() Shoe {
	return &adidasShoe{
		sneaker: sneaker{
			logo: "adidasFactory",
			size: 14,
		},
	}
}

func (a *adidasFactory) makeShirt() Clothing {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidasFactory",
			size: 14,
		},
	}
}

// nikeFactory is another Concrete Factory, it implements the abstract factory interface.
type nikeFactory struct{}

func (n *nikeFactory) makeShoe() Shoe {
	return &nikeShoe{
		sneaker: sneaker{
			logo: "nikeFactory",
			size: 14,
		},
	}
}

func (n *nikeFactory) makeShirt() Clothing {
	return &nikeShirt{
		shirt: shirt{
			logo: "nikeFactory",
			size: 14,
		},
	}
}

// Shoe is the interface of the first product.
type Shoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type sneaker struct {
	logo string
	size int
}

func (s *sneaker) setLogo(logo string) {
	s.logo = logo
}

func (s *sneaker) getLogo() string {
	return s.logo
}

func (s *sneaker) setSize(size int) {
	s.size = size
}

func (s *sneaker) getSize() int {
	return s.size
}

type adidasShoe struct {
	sneaker
}

type nikeShoe struct {
	sneaker
}

// Clothing is the interface of the second product.
type Clothing interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) getSize() int {
	return s.size
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}
