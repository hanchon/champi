package storecore

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func handleBytes(raw []byte) string {
	return hexutil.Encode(raw)
}

func handleUint(raw []byte) string {
	return new(big.Int).SetBytes(raw).String()
}

func handleInt(raw []byte, schemaType SchemaType) string {
	value := new(big.Int).SetBytes(raw)
	pow := int64(((schemaType - 32 + 1) * 8) - 1)
	compareValue := big.NewInt(2).Exp(big.NewInt(2), big.NewInt(pow), big.NewInt(0))
	if value.Cmp(compareValue) == 1 {
		offset := compareValue.Mul(compareValue, big.NewInt(-2))
		return value.Add(value, offset).String()
	}
	return value.String()
}

func handleBool(raw byte) bool {
	return raw == 1
}

func handleAddress(raw []byte) string {
	return common.BytesToAddress(raw).String()
}

func handleString(raw []byte) string {
	return string(raw)
}
