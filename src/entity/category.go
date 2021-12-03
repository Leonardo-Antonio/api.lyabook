package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" validmor:"required"`
	Slug      string             `bson:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty" validmor:"required"`
	Ean       string             `bson:"ean,omitempty" json:"ean,omitempty" xml:"ean,omitempty" validmor:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active    bool               `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
}

type Categories []Category
