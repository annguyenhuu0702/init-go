package models

import "time"

type Customer struct {
	Id          string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	Email       string    `json:"email" gorm:"column:email"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}
