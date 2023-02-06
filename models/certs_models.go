package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cert struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"name"`
	Platform string `json:"platform"`
}