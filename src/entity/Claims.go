package entity

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Email string `bson:"email" json:"email" xml:"email"`
	Name string `bson:"name" json:"name" xml:"name"`
	LastName string `bson:"last_name" json:"last_name" xml:"last_name"`
	Nick string `bson:"nick" json:"nick" xml:"nick"`
	jwt.StandardClaims
}
