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
