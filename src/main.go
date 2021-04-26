package main

import (
	"fmt"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/dbutil"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil { log.Fatalln(err) }
	db := dbutil.Connection()
	user := model.NewUser(db)
	id, _ := primitive.ObjectIDFromHex("6086022a3d7d45ae6640cfb8")
	res, _ := user.FindOne(id)
	fmt.Println(res.Name)
}