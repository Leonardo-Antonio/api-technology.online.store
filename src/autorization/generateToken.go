package autorization

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(user entity.User) (string, error) {
	claim := entity.Claim{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		Nick:     user.Nick,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			Issuer:    "Leonardo-Antonio",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
