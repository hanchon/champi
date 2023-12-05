package storecore

import "github.com/ethereum/go-ethereum/common/hexutil"

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func StringTo32Byte(s string) [32]byte {
	bytes, err := hexutil.Decode(s)
	if err != nil {
		panic("error decoding the hex string")
	}
	if len(bytes) != 32 {
		panic("invalid lenght")
	}
	return [32]byte(bytes)
}

func StringToByteArray(s string) []byte {
	bytes, err := hexutil.Decode(s)
	if err != nil {
		panic("error decoding the hex string")
	}
	return bytes
}
