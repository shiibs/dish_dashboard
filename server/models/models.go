package models

import "gorm.io/gorm"

// Dish represents a dish model with fields for dish details.
type Dish struct {
	gorm.Model
	ID          uint   `json:"dishId"`
	DishName    string `json:"dishName"`
	ImageURL    string `json:"imageUrl"`
	IsPublished bool   `json:"isPublished"`
}
