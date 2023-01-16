package models

type Product struct {
	Id    string  `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}
