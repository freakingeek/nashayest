package models

type Word struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
