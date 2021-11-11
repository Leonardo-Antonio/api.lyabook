package entity

type BookPayment struct {
	SoldBook SoldBook `bson:"sold_book,omitempty" json:"sold_book,omitempty" xml:"sold_book,omitempty"`
	Data     Book     `bson:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}
