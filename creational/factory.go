package creational

import "fmt"

//Factory Method is a creational design pattern that provides an interface for creating objects in a superclass,
//but allows subclasses to alter the type of objects that will be created.

// In this example, we are creating a factory method that creates a new objects that implement the Conveyor interface.

// Conveyor is the Product, the product is the one that declares the interface, which is common to all objects that
// can be produced by the creator and its subclasses.
type Conveyor interface {
	setName(name string)
	setSpeed(speed int)
	getName() string
	getSpeed() int
	transport() bool
}

type conveyance struct {
	name  string
	speed int
}

func (g *conveyance) setName(name string) {
	g.name = name
}

func (g *conveyance) getName() string {
	return g.name
}

func (g *conveyance) setSpeed(speed int) {
	g.speed = speed
}

func (g *conveyance) getSpeed() int {
	return g.speed
}

// truck is a concrete product. Concrete Products are different implementations of the product interface. Here we
// are using composition here just to simplify the product implementation.
type truck struct {
	conveyance
}

func (t truck) transport() bool {
	fmt.Printf("Truck %s is transporting goods at %d km/h\n", t.getName(), t.getSpeed())
	return true
}

func newTruck() Conveyor {
	return &truck{
		conveyance: conveyance{
			name:  "Truck conveyance",
			speed: 4,
		},
	}
}

// ship is a concrete product. Concrete Products are different implementations of the product interface.
type ship struct {
	conveyance
}

func (s ship) transport() bool {
	fmt.Printf("Ship %s is moving at %d knots\n", s.getName(), s.getSpeed())
	return true
}

func newShip() Conveyor {
	return &ship{
		conveyance: conveyance{
			name:  "Ship conveyance",
			speed: 1,
		},
	}
}

// airplane is a concrete product. Concrete Products are different implementations of the product interface.
type airplane struct {
	conveyance
}

func (a airplane) transport() bool {
	fmt.Printf("Airplane %s is flying at %d knots\n", a.getName(), a.getSpeed())
	return true
}

func newAirplane() Conveyor {
	return &airplane{
		conveyance: conveyance{
			name:  "Airplane conveyance",
			speed: 1,
		},
	}
}

// GetConveyance is the Factory function (or factory method) that returns an instance of the product.
func GetConveyance(conveyanceType string) (Conveyor, error) {
	switch conveyanceType {
	case "truck":
		return newTruck(), nil
	case "ship":
		return newShip(), nil
	case "airplane":
		return newAirplane(), nil
	default:
		return nil, fmt.Errorf("wrong conveyance type passed")
	}
}
