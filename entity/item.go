// package entites holds all entities shared across diff domain
package entity

import "github.com/google/uuid"

// An Item is an entity that represents a item in all domains
type Item struct {
	// Unique Identifier for an entity
	ID          uuid.UUID
	Name        string
	Description string
}
