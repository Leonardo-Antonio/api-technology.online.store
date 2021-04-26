package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sale struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	ProductsIDs []primitive.ObjectID `bson:"products_i_ds,omitempty" json:"products_ids,omitempty"`
	CreatedAt primitive.ObjectID `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.ObjectID `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt primitive.ObjectID `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}