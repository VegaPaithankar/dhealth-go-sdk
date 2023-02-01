package account

import (
	"dhealth-sdk/core/crypto/keypair"
	"dhealth-sdk/model/account/address"
	"dhealth-sdk/model/network/networktype"
)

type Account struct {
	addr    *address.Address
	keyPair *keypair.KeyPair
}

func CreateFromPrivateKey(privateKey string, networkType networktype.NetworkType) (*Account, error) {
	keyPair, err := keypair.CreateKeyPairFromPrivateKeyString(privateKey)
	if err != nil {
		return nil, err
	}

	pubKeyString, err := keypair.PublicKeyToString(keyPair.PublicKey())
	if err != nil {
		return nil, err
	}
	addr, err := address.CreateFromPublicKey(pubKeyString, networkType)
	if err != nil {
		return nil, err
	}
	return &Account{addr, keyPair}, nil
}
