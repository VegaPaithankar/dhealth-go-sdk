package keypair

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
)

var InvalidKeySizeError error = errors.New("invalid key size")

type KeyPair struct {
	privateKey []byte
	publicKey  []byte
}

func (kp KeyPair) PrivateKey() []byte {
	return kp.privateKey
}

func (kp KeyPair) PublicKey() []byte {
	return kp.publicKey
}

func CreateKeyPairFromPrivateKeyString(privateKeyString string) (*KeyPair, error) {
	hex_length := len(privateKeyString)
	if hex_length != 2*ed25519.SeedSize {
		return nil, InvalidKeySizeError
	}
	privateKeySeed := make([]uint8, ed25519.SeedSize)
	seed_length, err := hex.Decode(privateKeySeed, []byte(privateKeyString))
	if err != nil || seed_length != ed25519.SeedSize {
		return nil, InvalidKeySizeError
	}
	publicKey := ed25519.NewKeyFromSeed(privateKeySeed).Public()
	return &KeyPair{privateKeySeed, []byte(publicKey.(ed25519.PublicKey))}, nil
}

func PublicKeyToString(publicKey []byte) (string, error) {
	return hex.EncodeToString(publicKey), nil
}

func StringToPublicKey(publicKeyString string) ([]byte, error) {
	hex_length := len(publicKeyString)
	if hex_length != 2*ed25519.PublicKeySize {
		return nil, InvalidKeySizeError
	}
	publicKey := make([]uint8, ed25519.PublicKeySize)
	n, err := hex.Decode(publicKey, []byte(publicKeyString))
	if err != nil {
		return nil, err
	}
	if n != ed25519.PublicKeySize {
		return nil, InvalidKeySizeError
	}
	return publicKey, nil
}