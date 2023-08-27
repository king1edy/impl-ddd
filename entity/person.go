// package entites holds all entities shared across diff domain
package entity

import "github.com/google/uuid"

// A person is an entity that represents a person in all domains
type Person struct {
	// Unique Identifier for an entity
	ID   uuid.UUID
	Name string
	Age  int
}
