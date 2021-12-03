package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Commentary struct {
	IdClient primitive.ObjectID `bson:"id_client,omitempty" json:"id_client,omitempty" xml:"id_client,omitempty"`
	Message  string             `bson:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	Stars    uint8              `bson:"stars,omitempty" json:"stars,omitempty" xml:"stars,omitempty"`
}
