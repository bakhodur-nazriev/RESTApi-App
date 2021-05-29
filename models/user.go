package models

type Users struct {
	Id        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
	Birthday  string `json:"birthday"`
}
