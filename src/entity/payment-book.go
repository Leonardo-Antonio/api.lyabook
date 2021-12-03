package entity

type BookPayment struct {
	SoldBook  SoldBook `bson:"sold_book,omitempty" json:"sold_book,omitempty" xml:"sold_book,omitempty"`
	Data      Book     `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Tag       string   `bson:"tag,omitempty" json:"tag,omitempty" xml:"tag,omitempty"`
	CreatedAt string   `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
}
