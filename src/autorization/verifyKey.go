package autorization

import (
	"errors"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"github.com/dgrijalva/jwt-go"
)

func validateToken(t string) (entity.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &entity.Claim{}, verifyFunction)
	if err != nil {
		return entity.Claim{}, err
	}
	if !token.Valid {
		return entity.Claim{}, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*entity.Claim)
	if !ok {
		return entity.Claim{}, errors.New("could not get claims")
	}
	return *claim, nil
}

func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
