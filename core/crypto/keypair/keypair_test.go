package keypair

import (
	"bytes"
	"testing"
)

func TestCreateKeyPairFromPrivateKeyString(t *testing.T) {
	privateKeyString := "059121A6A6A49D0A42DBC5FD02DED32F6DC3F62EC94DE4EEDF724E261F1D3678"
	keyPair, err := CreateKeyPairFromPrivateKeyString(privateKeyString)
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte{
		0x94, 0xA0, 0x7B, 0x04, 0xFE, 0xD4, 0x62, 0x8F,
		0x40, 0xE1, 0x95, 0xCF, 0xB5, 0x36, 0x74, 0xC2,
		0x15, 0xAE, 0x9E, 0x5E, 0x82, 0x43, 0xCB, 0xA1,
		0xA2, 0x3F, 0x83, 0x8C, 0x54, 0x5B, 0xB0, 0xF3,
	}
	pubKey := keyPair.PublicKey()
	if bytes.Compare([]uint8(pubKey), expected) != 0 {
		t.Fatalf("computed public key %v does not match expected %v", keyPair.PublicKey(), expected)
	}
}
