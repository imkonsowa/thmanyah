package keys

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewKeyStore,
)

type Store struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewKeyStore() *Store {
	privateKey, err := readPrivateKey()
	if err != nil {
		panic(fmt.Errorf("getting private key: %w", err))
	}

	publicKey, err := readPublicKey()
	if err != nil {
		panic(fmt.Errorf("getting public key: %w", err))
	}

	return &Store{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (s *Store) PrivateKey() *rsa.PrivateKey {
	return s.privateKey
}

func (s *Store) PublicKey() *rsa.PublicKey {
	return s.publicKey
}

func readPublicKey() (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile("../../keys/public_key.pem")
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func readPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile("../../keys/private_key.pem")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing auth private key: %w", err)
	}

	return privateKey, nil
}
