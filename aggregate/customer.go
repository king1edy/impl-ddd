// package aggregate holds our aagrets that combines many entities into a full obejct
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/king1edy/impl-ddd/entity"
	"github.com/king1edy/impl-ddd/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	// Person is the root entity of a customer.
	// which means person.ID is the main identifier for the customer
	person   *entity.Person
	products []*entity.Item

	transactions []valueobject.Transaction
}

// The NewCustomer is a factory to create a new customer whenever need be.
// it will validate that the name of the cusomer is not empty.
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
