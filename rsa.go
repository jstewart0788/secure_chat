package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return privateKey, &privateKey.PublicKey
}

func publicKeyToString(p *rsa.PublicKey) string {
	pubString, err := x509.MarshalPKIXPublicKey(p)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	pubPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubString,
		},
	)

	return string(pubPem)
}

func privateKeyToString(p *rsa.PrivateKey) string {
	secretString := x509.MarshalPKCS1PrivateKey(p)
	secretPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: secretString,
		},
	)
	return string(secretPem)
}

func decodePubKey(pubString string) (*rsa.PublicKey, error) {
	decodedPub, _ := pem.Decode([]byte(pubString))
	if decodedPub == nil {
		return nil, errors.New("Error Parsing Public Key")
	}
	pub, err2 := x509.ParsePKIXPublicKey(decodedPub.Bytes)
	if err2 != nil {
		fmt.Println("Error2: ", err2)
		os.Exit(1)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
	}
	return nil, errors.New("Error Parsing Public Key")
}

func decodeSecretKey(secretString string) (*rsa.PrivateKey, error) {
	decodedSecret, _ := pem.Decode([]byte(secretString))
	if decodedSecret == nil {
		return nil, errors.New("Error Parsing Public Key")
	}
	secret, err2 := x509.ParsePKCS1PrivateKey(decodedSecret.Bytes)
	if err2 != nil {
		fmt.Println("Error2: ", err2)
		os.Exit(1)
	}
	return secret, nil

}

func getPublicRSAKeyFromFile() string {
	pubKeyString, err := ioutil.ReadFile("../Shared/keys/rsa/public.pem")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return string(pubKeyString)
}

func getPrivateRSAKeyFromFile() string {
	pubKeyString, err := ioutil.ReadFile("./keys/rsa/private.pem")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return string(pubKeyString)
}

func getEncryptedAESKeyFromFile() *big.Int {
	var c big.Int
	cypher, err := ioutil.ReadFile("../Shared/cypher/aes/key")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	c.SetString(string(cypher), 10)
	return &c
}

func getIVFromFile() string {

	iv, err := ioutil.ReadFile("../Shared/keys/aes/iv")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return string(iv)
}

func releasePublicRSAKey(s string) error {
	return ioutil.WriteFile("../Shared/keys/rsa/public.pem", []byte(s), 0666)
}

func storePrivateRSAKey(s string) error {
	return ioutil.WriteFile("./keys/rsa/private.pem", []byte(s), 0666)
}

func releaseEncryptedAESKey(c *big.Int) error {
	return ioutil.WriteFile("../Shared/cypher/aes/key", []byte(c.Text(10)), 0666)
}

func releaseIV(iv uint64) error {
	return ioutil.WriteFile("../Shared/keys/aes/iv", []byte(strconv.FormatUint(iv, 10)), 0666)
}

func encryptRSA(p *rsa.PublicKey, message uint64) *big.Int {
	var c, e, m big.Int

	e.SetInt64(int64(p.E))
	m.SetUint64(message)
	c.Exp(&m, &e, p.N)
	return &c
}

func decrptRSA(p *rsa.PrivateKey, c *big.Int) uint64 {
	var m big.Int
	m.Exp(c, p.D, p.N)
	message := m.Uint64()
	return message
}
