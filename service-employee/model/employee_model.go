package model

type Employee struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
