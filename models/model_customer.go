package models

type Customer struct {
	CustomerID string `json:"customerId" gorm:"primaryKey"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}
