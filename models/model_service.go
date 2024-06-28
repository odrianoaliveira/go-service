package models

type Service struct {
	ServiceID string  `json:"serviceId" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
