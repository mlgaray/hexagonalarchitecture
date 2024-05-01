package models

type Incident struct {
	Category string  `json:"category,omitempty" bson:"category,omitempty"`
	User     []*User `json:"user,omitempty" bson:"user,omitempty"`
}
