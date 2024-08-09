package dtos

type CreateCustomerDTO struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"omitempty"`
}

type UpdateCustomerDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"omitempty"`
}
