package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object becuase it has no identifier
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
