package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Categories struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedBy primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	CreatedAt primitive.ObjectID `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt primitive.ObjectID `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt primitive.ObjectID `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`

}