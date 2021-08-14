package utilities

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

type crypto struct{}

var Crypto *crypto

func init() {
	once.Do(func() {
		Crypto = new(crypto)
	})
}

func (*crypto) KeyGen() (k *ecdsa.PrivateKey, err error) {
	if k, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		return nil, err
	}
	return k, nil
}

func (*crypto) PemKeyPair(key *ecdsa.PrivateKey) (privateKeyPEM []byte, publicKeyPEM []byte, err error) {
	var der []byte
	if der, err = x509.MarshalECPrivateKey(key); err != nil {
		return nil, nil, err
	}

	privateKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	})

	if der, err = x509.MarshalPKIXPublicKey(key.Public()); err != nil {
		return nil, nil, err
	}

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: der,
	})
	return
}
