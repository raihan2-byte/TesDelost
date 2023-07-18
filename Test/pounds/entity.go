package pounds

import (
	"delos/farm"

	"gorm.io/gorm"
)

type Pounds struct {
	// ID        int
	gorm.Model
	Name      string
	FarmID int 
	Farm farm.Farm 
	// CreatedAt time.Time
	// UpdatedAt time.Time
}