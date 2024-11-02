package models

type Todo struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Item       string `json:"item"`
	IsComplete bool   `json:"isComplete"`
}