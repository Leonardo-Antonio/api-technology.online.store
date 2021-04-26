package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	Name string `bson:"name,omitempty," json:"name,omitempty"`
	LastName string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Rol string `bson:"rol,omitempty" json:"rol,omitempty"`
	Nick string `bson:"nick,omitempty" json:"nick,omitempty"`
	Image string `bson:"image,omitempty" json:"image,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedBy primitive.ObjectID `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedBy primitive.ObjectID `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty"`
}
