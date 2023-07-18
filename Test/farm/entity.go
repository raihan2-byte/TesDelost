package farm

import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	Name      string
}
