package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	Name         string               `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" validmor:"required"`
	Slug         string               `bson:"slug,omitempty" json:"slug,omitempty" xml:"slug,omitempty"`
	Author       string               `bson:"author,omitempty" json:"author,omitempty" xml:"author,omitempty" validmor:"required"`
	Editorial    string               `bson:"editorial,omitempty" json:"editorial,omitempty" xml:"editorial,omitempty" validmor:"required"`
	PriceCurrent float64              `bson:"price_current,omitempty" json:"price_current,omitempty" xml:"price_current,omitempty" validmor:"required"`
	PriceBefore  float64              `bson:"price_before,omitempty" json:"price_before,omitempty" xml:"price_before,omitempty"`
	Stars        uint                 `bson:"stars,omitempty" json:"stars,omitempty" xml:"stars,omitempty"`
	Description  string               `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty" validmor:"required"`
	Commentaries []Commentary         `bson:"commentaries,omitempty" json:"commentaries,omitempty" xml:"commentaries,omitempty"`
	Type         Format               `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Details      []string             `bson:"details,omitempty" json:"details,omitempty" xml:"details,omitempty" validmor:"required"`
	Categories   []primitive.ObjectID `bson:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty" validmor:"required"`
	ImagesSrc    []string             `bson:"images_src,omitempty" json:"images_src,omitempty" xml:"images_src,omitempty" validmor:"required"`
	CreatedAt    time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt    time.Time            `bson:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt    time.Time            `bson:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active       bool                 `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
}

type Books []Book
