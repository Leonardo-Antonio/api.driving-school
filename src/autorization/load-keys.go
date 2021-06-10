package autorization

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	once      sync.Once
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func LoadFiles() error {
	var err error
	once.Do(func() {
		err = loadFiles()
	})
	return err
}

func loadFiles() error {
	privateByte, err := ioutil.ReadFile("./src/autorization/keys/app.rsa")
	if err != nil {
		return err
	}

	publicByte, err := ioutil.ReadFile("./src/autorization/keys/app.rsa.pub")
	if err != nil {
		return err
	}

	return parseRSA(privateByte, publicByte)
}

func parseRSA(privateByte, publicByte []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if err != nil {
		return err
	}
	return nil
}
