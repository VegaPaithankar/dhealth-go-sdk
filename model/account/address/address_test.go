package address

import (
	"log"
	"testing"

	"dhealth-sdk/model/network/networktype"
)

func TestCreateFromPublicKey(t *testing.T) {
	type testcase struct {
		pubkey        string
		expected_addr string
	}
	testcases := []testcase{
		{"94A07B04FED4628F40E195CFB53674C215AE9E5E8243CBA1A23F838C545BB0F3",
			"NDO5BQDXE2CVIT32F72V4FF3IAEVYBS2SYMLVOQ"},
		{"566067A835D5F3328C532F6AFEE7DE87861F73685A6D9B2B47450BE3E507727B",
			"NCGHO7XQPZUE7LDJD2P4HYFZNSL4CN27CPQKZWQ"},
	}
	for _, tc := range testcases {
		log.Printf("pubkey length=%d", len(tc.pubkey))
		addr, err := CreateFromPublicKey(tc.pubkey, networktype.MAIN_NET)
		if err != nil {
			t.Fatal(err)
		}
		if addr.Address != tc.expected_addr {
			t.Fatalf("computed address %s does not match expected %s", addr.Address, tc.expected_addr)
		}
		if addr.NetworkType != networktype.MAIN_NET {
			t.Fatalf("computed networktype %v does not match expected %v", addr.NetworkType, networktype.MAIN_NET)
		}
	}
}
