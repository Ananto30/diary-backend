package repository

import "github.com/jinzhu/gorm"

type PgRepo struct {
	db *gorm.DB
}

