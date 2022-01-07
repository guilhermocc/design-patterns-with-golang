package creational

// Abstract Factory is a creational design pattern that lets you produce families of related objects without
// specifying their concrete classes.

import "fmt"

type SportsFactory interface {
	makeShoe() Shoe
	makeShirt() Shirt
}

func getSportsFactory(brand string) (SportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}

type adidas struct{}

func (a *adidas) makeShoe() Shoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShirt() Shirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

type nike struct{}

func (n *nike) makeShoe() Shoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShirt() Shirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}

type Shoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

type Shirt interface {
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
