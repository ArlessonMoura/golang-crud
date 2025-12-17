package common

import (
	"gorm.io/gorm"
)

type Dependencies struct {
	DB *gorm.DB
}

func (d *Dependencies) Load() error {
	return nil
}
