package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Text  string
}
