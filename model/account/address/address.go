package address

import (
	"encoding/hex"
	"errors"
	"log"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"

	"dhealth-sdk/core/format/base32"
	"dhealth-sdk/model/network/networktype"
)

const PUBLIC_KEY_SIZE_BYTES = 32
const ADDRESS_DECODED_SIZE_BYTES = 24
const ADDRESS_ENCODED_SIZE_BYTES = 39 // 24 bytes = 192 bits. 5 bits / encoded char => ceil(192/5) = 39
const ADDRESS_CHECKSUM_SIZE_BYTES = 3

var InvalidKeySizeError error = errors.New("invalid key size")

type Address struct {
	Address     string
	NetworkType networktype.NetworkType
}

func CreateFromPublicKey(publicKeyString string, networkType networktype.NetworkType) (*Address, error) {
	length := len(publicKeyString) / 2
	if length != PUBLIC_KEY_SIZE_BYTES {
		log.Printf("pubkey string length %d is not 2x pubkey size (%d)", len(publicKeyString), PUBLIC_KEY_SIZE_BYTES)
		return nil, InvalidKeySizeError
	}
	publicKey := make([]byte, length)
	n, err := hex.Decode(publicKey, []byte(publicKeyString))
	if err != nil {
		log.Printf("hex decode error: %s", publicKeyString)
		return nil, err
	}
	if n != length {
		log.Printf("extracted pubkey size %d instead of %d", n, length)
		return nil, InvalidKeySizeError
	}
	addr := publicKeyToAddress(publicKey, networkType)
	addrString, err := addressToString(addr)
	if err != nil {
		return nil, err
	}
	return &Address{Address: addrString, NetworkType: networkType}, nil
}

func publicKeyToAddress(publicKey []byte, networkType networktype.NetworkType) []byte {
	// step 1: sha3 hash of the public key
	publicKeyHash := sha3.Sum256(publicKey)

	// step 2: ripemd160 hash of (1)
	hash := ripemd160.New()
	hash.Write(publicKeyHash[:])
	ripemdHash := hash.Sum(nil)
	log.Printf("step 2: size of ripemd hash=%d", len(ripemdHash))

	// step 3: add network identifier byte in front
	decodedAddress := make([]byte, 0, ADDRESS_DECODED_SIZE_BYTES)
	decodedAddress = append(decodedAddress, byte(networkType))
	log.Printf("step 3: addr bytes=%v", decodedAddress)

	// step 4: append (2)
	decodedAddress = append(decodedAddress, ripemdHash...)
	log.Printf("step 4: addr bytes=%v", decodedAddress)

	// step 5: concatenate (3) and the checksum of (3)
	checksum := sha3.Sum256(decodedAddress)

	// step5: append first <checksum_size> bytes of checksum
	decodedAddress = append(decodedAddress, checksum[0:ADDRESS_CHECKSUM_SIZE_BYTES]...)
	log.Printf("step 5: addr bytes=%v", decodedAddress)

	return decodedAddress
}

func addressToString(addr []byte) (string, error) {
	if len(addr) != ADDRESS_DECODED_SIZE_BYTES {
		log.Printf("raw address size %d does not match expected (%d)", len(addr), ADDRESS_DECODED_SIZE_BYTES)
		return "", InvalidKeySizeError
	}
	// return string(base32.Encode(addr)[0:ADDRESS_ENCODED_SIZE_BYTES]), nil
	return string(base32.Encode(addr)[:]), nil
}
