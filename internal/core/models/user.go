package models

type User struct {
	ID    int    `json:"id,omitempty" json:"id"`
	Name  string `json:"name,omitempty" json:"name" bson:"name,omitempty"`
	Email string `json:"email,omitempty" json:"email"  bson:"email,omitempty"`
}
