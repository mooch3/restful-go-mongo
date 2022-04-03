package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	Name string `json:"name,omtiempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title string `json:"title,omitempty" validate:"required"`
}