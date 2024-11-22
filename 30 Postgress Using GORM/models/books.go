package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}