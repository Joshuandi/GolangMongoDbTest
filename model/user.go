package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Username string             `json:"username,omitempty" validate: "required"`
	Email    string             `json:"email,omitempty" validate: "required"`
	Password string             `json:"password,omitempty" validate: "required"`
	Age      int                `json:"age,omitempty" validate: "required"`
}

type UserGetAll struct {
	Idusers  primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
}

type UserUpdateInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
