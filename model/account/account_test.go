package account

import (
	"log"
	"testing"

	"dhealth-sdk/model/network/networktype"
)

func TestCreateFromPrivateKey(t *testing.T) {
	type testcase struct {
		privateKey    string
		expected_addr string
	}
	testcases := []testcase{
		{"059121A6A6A49D0A42DBC5FD02DED32F6DC3F62EC94DE4EEDF724E261F1D3678",
			"NDO5BQDXE2CVIT32F72V4FF3IAEVYBS2SYMLVOQ"},
	}
	for _, tc := range testcases {
		log.Printf("privateKey length=%d", len(tc.privateKey))
		account, err := CreateFromPrivateKey(tc.privateKey, networktype.MAIN_NET)
		if err != nil {
			t.Fatal(err)
		}
		if account.addr.Address != tc.expected_addr {
			t.Fatalf("computed address %s does not match expected %s", account.addr.Address, tc.expected_addr)
		}
		if account.addr.NetworkType != networktype.MAIN_NET {
			t.Fatalf("computed networktype %v does not match expected %v", account.addr.NetworkType, networktype.MAIN_NET)
		}
	}
}
