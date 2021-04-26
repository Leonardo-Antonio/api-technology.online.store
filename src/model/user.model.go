package model

import (
	"context"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type user struct {
	db *mongo.Database
	collection *mongo.Collection
}

func NewUser(db *mongo.Database) *user {
	return &user{db, db.Collection("users")}
}

func (u *user) Create(userEntity entity.User) (result *mongo.InsertOneResult, err error) {
	userEntity.CreatedAt = time.Now()
	userEntity.Active = true
	result, err = u.collection.InsertOne(context.TODO(), &userEntity)
	if err != nil {
		return
	}
	return result, nil
}

func (u *user) FindAll() (users []entity.User, err error) {
	var cursor *mongo.Cursor
	filter := bson.M{
		"active": true,
	}
	cursor, err = u.collection.Find(context.TODO(), filter)
	if err != nil {
		return
	}
	if err := cursor.All(context.TODO(), &users); err != nil {
		return users, err
	}
	return
}

func (u *user) FindOne(ID primitive.ObjectID) (userEntity entity.User, err error) {
	filter := bson.M{
		"_id": ID,
		"active": true,
	}
	err = u.collection.FindOne(context.TODO(), filter).Decode(&userEntity)
	if err != nil {
		return
	}
	return
}

func (u *user) UpdateByID(userEntity entity.User) (result *mongo.UpdateResult, err error) {
	userEntity.UpdatedAt = time.Now()
	update := bson.M{
		"$set":
			userEntity,
	}
	result, err = u.collection.UpdateByID(context.TODO(), userEntity.ID, update)
	if err != nil {
		return
	}
	if result.MatchedCount != 1 {
		return result, utils.ErrUpdateOne
	}
	return
}

func (u *user) RemoveByID(ID primitive.ObjectID) (result *mongo.UpdateResult, err error) {
	remove := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}
	result, err = u.collection.UpdateByID(context.TODO(), ID, remove)
	if err != nil {
		return
	}
	return
}