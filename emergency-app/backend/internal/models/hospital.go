package models

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Blood   int    `json:"blood_units"`
	Oxygen  int    `json:"oxygen_units"`
	Medicine string `json:"medicine_stock"`
}
