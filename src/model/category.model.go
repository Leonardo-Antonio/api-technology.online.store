package model

import (
	"context"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type (
	category struct {
		db         *mongo.Database
		collection *mongo.Collection
	}
	ICategory interface {
		Create(categoryEntity entity.Category) (result *mongo.InsertOneResult, err error)
		FindAll() (categories []entity.Category, err error)
		FindOne(ID primitive.ObjectID) (entityCategory entity.Category, err error)
		Update(entityCategory entity.Category) (result *mongo.UpdateResult, err error)
		Delete(ID primitive.ObjectID) (result *mongo.UpdateResult, err error)
	}
)

func NewCategory(db *mongo.Database) *category {
	return &category{
		db:         db,
		collection: db.Collection("categories"),
	}
}
func (c *category) Create(categoryEntity entity.Category) (result *mongo.InsertOneResult, err error) {
	categoryEntity.ID = primitive.NewObjectID()
	categoryEntity.Name = strings.Title(categoryEntity.Name)
	categoryEntity.CreatedAt = time.Now()
	categoryEntity.Active = true
	result, err = c.collection.InsertOne(context.TODO(), &categoryEntity)
	if err != nil {
		return nil, err
	}
	return
}

func (c category) FindAll() (categories []entity.Category, err error) {
	filter := bson.M{
		"active": true,
	}
	cursor, err := c.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if cursor.All(context.TODO(), &categories) != nil {
		return nil, err
	}
	return
}

func (c *category) FindOne(ID primitive.ObjectID) (entityCategory entity.Category, err error) {
	filter := bson.M{
		"_id":    ID,
		"active": true,
	}
	if c.collection.FindOne(context.TODO(), filter).Decode(&entityCategory) != nil {
		return
	}
	return
}

func (c *category) Update(entityCategory entity.Category) (result *mongo.UpdateResult, err error) {
	entityCategory.UpdatedAt = time.Now()
	update := bson.M{
		"$set": entityCategory,
	}
	result, err = c.collection.UpdateByID(context.TODO(), entityCategory.ID, update)
	if err != nil {
		return
	}
	return
}

func (c *category) Delete(ID primitive.ObjectID) (result *mongo.UpdateResult, err error) {
	remove := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}
	result, err = c.collection.UpdateByID(context.TODO(), ID, remove)
	if err != nil {
		return
	}
	return
}
