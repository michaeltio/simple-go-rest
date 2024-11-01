package models

type Todo struct {
	ID         int    `json:"id"`
	Item       string `json:"item"`
	IsComplete bool   `json:"isComplete"`
}
