package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID           primitive.ObjectID   `bson:"_id" json:"_id"`
	Name         string               `bson:"name,omitempty" json:"name,omitempty"`
	CategoriesID []primitive.ObjectID `bson:"categories_id,omitempty" json:"categories_id,omitempty"`
	Labels       []string             `bson:"labels,omitempty" json:"labels,omitempty"`
	Price        float32              `bson:"price,omitempty" json:"price,omitempty"`
	Description  string               `bson:"description,omitempty" json:"description,omitempty"`
	Stock        uint                 `bson:"stock,omitempty" json:"stock,omitempty"`
	CreatedBy    primitive.ObjectID   `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy    primitive.ObjectID   `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedBy    primitive.ObjectID   `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	CreatedAt    primitive.ObjectID   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt    primitive.ObjectID   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt    primitive.ObjectID   `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
