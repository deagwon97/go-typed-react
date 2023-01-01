package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}
