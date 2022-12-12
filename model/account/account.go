package account

import (
	"dhealth-sdk/core/crypto/keypair"
	"dhealth-sdk/model/network/networktype"
)

type Account struct {
	addr    Address
	keyPair keypair.KeyPair
}

func (account *Account) CreateFromPrivateKey(privateKey string, networkType networktype.NetworkType) {
	keyPair := keypair.CreateKeyPairFromPrivateKeyString(privateKey)
	addr := rawaddress.AddressToString(rawAddress.PublieKeyToAddress(keyPair.PublicKey(), networkType))
	return &Account{addr: address.CreateFromRawAddress(addr), keyPair}
}
