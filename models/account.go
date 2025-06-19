package models

type Account struct {
	ID         uint    `json:"id" gorm:"primary"`
	CustomerID string  `json:"customer_id"`
	Type       string  `json:"type"`
	Balance    float64 `json:"balance"`
	Status     string  `json:"status"`
}
