package keystore

import (
	"github.com/stretchr/testify/assert"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestParseKeyFromhex(t *testing.T) {
	testAddrPk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	testAddr := ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	key, err := ParseKeyFromHex(testAddrPk)
	assert.Nil(t, err)

	pubKeyAddress := crypto.PubkeyToAddress(key.PublicKey)
	assert.Equal(t, pubKeyAddress, testAddr)
}
