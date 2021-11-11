package entity

type SoldBook struct {
	Id        string `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	TimesSold uint   `bson:"times_sold,omitempty" json:"times_sold,omitempty" xml:"times_sold,omitempty"`
	BooksSold uint   `bson:"books_sold,omitempty" json:"books_sold,omitempty" xml:"books_sold,omitempty"`
}

type SoldBooks []*SoldBook
