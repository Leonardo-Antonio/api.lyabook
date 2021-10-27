package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ProductPayment struct {
		Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
		IdPayment   string             `bson:"id_client,omitempty" json:"id_client,omitempty" xml:"id_client,omitempty"`
		Title       string             `bson:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
		PriceUnit   float32            `bson:"unit_price,omitempty" json:"unit_price,omitempty" xml:"unit_price,omitempty"`
		Quantity    uint               `bson:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
		Description string             `bson:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
		PictureUrl  string             `bson:"picture_url,omitempty" json:"picture_url,omitempty" xml:"picture_url,omitempty"`
		CategoryId  string             `bson:"category_id,omitempty" json:"category_id,omitempty" xml:"category_id,omitempty"`
	}

	PaymentClient struct {
		Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
		IdClient  primitive.ObjectID `bson:"id_client,omitempty" json:"id_client,omitempty" xml:"id_client,omitempty"`
		IdPayment int                `bson:"payment_id,omitempty" json:"payment_id,omitempty" xml:"payment_id,omitempty"`
		Status    string             `bson:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
		Products  []ProductPayment   `bson:"products,omitempty" json:"products,omitempty" xml:"products,omitempty"`
		Client    User               `bson:"client,omitempty" json:"client,omitempty" xml:"client,omitempty"`
		CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" xml:"created_at,omitempty"`
		UpdateAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" xml:"updated_at,omitempty"`
		DeleteAt  time.Time          `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
		Active    bool               `json:"active,omitempty" bson:"active,omitempty" xml:"active,omitempty"`
	}
)
