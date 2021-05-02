package autorization

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

var (
	signKey *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func LoadKeys()  {
	// key private
	signKeyByte, err := ioutil.ReadFile("src/certificate/app.rsa")
	if err != nil {
		log.Fatalln(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	if err != nil {
		log.Fatalln(err)
	}

	// key public
	verifyKeyByte, err := ioutil.ReadFile("src/certificate/app.rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		log.Fatalln(err)
	}
}
