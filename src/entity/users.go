package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" xml:"_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty" xml:"name,omitempty"`
	LastName         string             `json:"last_name,omitempty" bson:"last_name,omitempty" xml:"last_name,omitempty"`
	Dni              string             `json:"dni,omitempty" bson:"dni,omitempty" xml:"dni,omitempty"`
	Email            string             `json:"email,omitempty" bson:"email,omitempty" xml:"email,omitempty"`
	VerificationCode string             `json:"verification_code,omitempty" bson:"verification_code,omitempty" xml:"verification_code,omitempty"`
	Password         string             `json:"password,omitempty" bson:"password,omitempty" xml:"password,omitempty" validmor:"string,min=8"`
	Rol              string             `json:"rol,omitempty" bson:"rol,omitempty" xml:"rol,omitempty"`
	CreateAt         time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdateAt         time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeleteAt         time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active           bool               `json:"active,omitempty" bson:"active,omitempty" xml:"active,omitempty"`
}

type Users []User
