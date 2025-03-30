package models

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Type        string `json:"type"`
	BloodType   string `json:"blood_type"`
	OxygenUnits int    `json:"oxygen_units"`
	Medicine    string `json:"medicine"`
	Status      string `json:"status"`
	HospitalID  uint   `json:"hospital_id"`
}
