package models

type Vendor struct {
	VendorID string `json:"vendorId" gorm:"primaryKey"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
