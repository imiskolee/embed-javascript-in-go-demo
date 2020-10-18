package main

import "github.com/jinzhu/gorm"

type CustomEventModel struct {
	gorm.Model
	ModelName string
	ClientID string
	Content string
	Enabled bool
}

