package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Mobile    string             `bson:"mobile"`
	Password  string             `bson:"password"`
	TwoFactor bool               `bson:"twoFactor"`
}
