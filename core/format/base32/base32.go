package base32

import "log"

const RAW_BLOCK_SIZE = 5
const BASE32_BLOCK_SIZE = 8
const BITS_PER_BYTE = 8

type encoding_spec struct {
	straddle          bool
	start_byte_offset int8
	upper_byte_shift  int8 // if straddle shift left, otherwise shift right by this amount.
	upper_byte_mask   byte
	lower_byte_shift  int8
}

var block_encoding_spec = []encoding_spec{
	{false, 0, 3, 0x1f, 0},
	{true, 0, 2, 0x7, 6},
	{false, 1, 1, 0x1f, 0},
	{true, 1, 4, 0x1, 4},
	{true, 2, 1, 0xf, 7},
	{false, 3, 2, 0x1f, 0},
	{true, 3, 3, 0x3, 5},
	{false, 4, 0, 0x1f, 0},
}

var BASE32DIGIT = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z', '2', '3', '4', '5', '6', '7',
}

func Encode(raw []byte) []byte {
	log.Printf("base32.Encode: raw=%v", raw)
	length := len(raw)
	output_length := (length*BITS_PER_BYTE + RAW_BLOCK_SIZE - 1) / RAW_BLOCK_SIZE
	padding := RAW_BLOCK_SIZE - length%RAW_BLOCK_SIZE
	if padding == RAW_BLOCK_SIZE {
		padding = 0
	}
	if padding != 0 {
		b := make([]byte, padding)
		raw = append(raw, b...)
	}
	b32 := make([]byte, 0, len(raw)/RAW_BLOCK_SIZE*BASE32_BLOCK_SIZE)
	for idx := 0; idx < len(raw); idx += RAW_BLOCK_SIZE {
		b32 = append(b32, encodeBlock(raw[idx:idx+RAW_BLOCK_SIZE])...)
	}
	return b32[:output_length]
}

func Decode(b32 []byte) []byte {
	return b32
}

func encodeBlock(raw []byte) []byte {
	log.Printf("base32.encodeBlock: raw=%v", raw)
	b32 := make([]byte, 0, BASE32_BLOCK_SIZE)
	for i := 0; i < BASE32_BLOCK_SIZE; i++ {
		rb := rawBits(raw, i)
		log.Printf("  %d: %d (%c)", i, rb, BASE32DIGIT[rb])
		b32 = append(b32, BASE32DIGIT[rawBits(raw, i)])
	}
	return b32
}

func decodeBlock(b32 []byte) []byte {
	return b32
}

func rawBits(raw []byte, n int) byte {
	bes := block_encoding_spec[n]
	if bes.straddle {
		return ((raw[bes.start_byte_offset] & bes.upper_byte_mask) << bes.upper_byte_shift) |
			(raw[bes.start_byte_offset+1] >> bes.lower_byte_shift)
	} else {
		return (raw[bes.start_byte_offset] >> bes.upper_byte_shift) & bes.upper_byte_mask
	}
}
