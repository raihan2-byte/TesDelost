package farm

import (
	"gorm.io/gorm"
)

type Farm struct {
	// ID        int
	gorm.Model
	Name      string
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt 
}
