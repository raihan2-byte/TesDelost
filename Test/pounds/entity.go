package pounds

import (
	"delos/farm"

	"gorm.io/gorm"
)

type Pounds struct {
	gorm.Model
	Name      string
	FarmID int 
	Farm farm.Farm 
	
}
