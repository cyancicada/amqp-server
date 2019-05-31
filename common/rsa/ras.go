package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

type (
	Rsa struct {
		publicKey  []byte
		privateKey []byte
	}
)

//1. yum install -y openssl
//2. openssl genrsa -out rsa_private_key.pem 10240
//3. openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func NewRsa(publicKeyPath, privateKeyPath string) (*Rsa, error) {
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return &Rsa{publicKey: publicKey, privateKey: privateKey}, nil
}

func (r *Rsa) Encrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(base64.StdEncoding.EncodeToString([]byte(origData))))
}

func (r *Rsa) Decrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	origin, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(origin))
}
