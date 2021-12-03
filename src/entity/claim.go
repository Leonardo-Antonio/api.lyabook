package entity

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaimUser struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" xml:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" xml:"name,omitempty"`
	LastName string             `json:"last_name,omitempty" bson:"last_name,omitempty" xml:"last_name,omitempty"`
	Dni      string             `json:"dni,omitempty" bson:"dni,omitempty" xml:"dni,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" xml:"email,omitempty"`
	Rol      string             `json:"rol,omitempty" bson:"rol,omitempty" xml:"rol,omitempty"`
	CreateAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdateAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeleteAt time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active   bool               `json:"active,omitempty" bson:"active,omitempty" xml:"active,omitempty"`
	jwt.StandardClaims
}
