package publisher

import (
	"context"
	"math/big"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/brianvoe/gofakeit/v7"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/keystore"

	ethsim "github.com/ethereum/go-ethereum/ethclient/simulated"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func simTestBackend(testAddr ethcommon.Address) *ethsim.Backend {
	balance, ok := big.NewInt(0).SetString("90000000000000000000", 10)
	if !ok {
		panic("Failed to parse bigInt")
	}
	return ethsim.NewBackend(
		ethtypes.GenesisAlloc{
			testAddr: {Balance: balance},
		},
	)
}

func TestSendUpdate(t *testing.T) {
	ctx := context.Background()
	// Create eth fake backend
	testAddrPk := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	testAddr := ethcommon.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	backend := simTestBackend(testAddr)
	defer backend.Close()

	client := backend.Client()
	chainID, err := client.ChainID(ctx)
	assert.Nil(t, err)
	key, err := keystore.ParseKeyFromHex(testAddrPk)
	assert.Nil(t, err)

	transactor, err := NewTransactor(ctx, client, key, chainID, testAddr)
	assert.Nil(t, err)

	var merkleUpdate MerkleRootUpdate
	err = gofakeit.Struct(&merkleUpdate)
	assert.Nil(t, err)


	// Temporarily assert error due to `no contract code at given address`
	err = transactor.SendUpdate(&merkleUpdate)
	assert.Error(t, err)
}
