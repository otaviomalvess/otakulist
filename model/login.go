package model

// Login ..
type Login struct {
	Email    string `bson:"email"    json:"email"    validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}
