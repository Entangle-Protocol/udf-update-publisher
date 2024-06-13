package publisher

import (
	"bytes"
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/config"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/tests/update"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/types"

	mockFetcher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	mockTransactor "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	// mockPublisher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
)

func TestPublishUpdate(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()
	fetcherMock := mockFetcher.NewMockIFetcher(t)

	dataKey := "NGL/USDT"

	// Generate test proofs and set mocker rv to return them
	key, err := crypto.GenerateKey()
	r.NoError(err)

	proof, err := update.GenerateProof(key, dataKey)
	r.NoError(err)

	fetcherMock.On("GetFeedProofs", ctx, dataKey).Return(proof, nil)

	transactorMock1 := mockTransactor.NewMockITransactor(t)
	transactorMock2 := mockTransactor.NewMockITransactor(t)

	merkleUpdate, err := NewMerkleUpdateFromProof(proof)
	r.NoError(err)

	transactorMock1.On("SendUpdate", merkleUpdate).Return(nil, nil)
	transactorMock1.On("LatestUpdate", merkleUpdate.DataKey).Return(big.NewInt(0), big.NewInt(0))
	transactorMock2.On("SendUpdate", merkleUpdate).Return(nil, nil)
	transactorMock2.On("LatestUpdate", merkleUpdate.DataKey).Return(big.NewInt(0), big.NewInt(0))

	conf := config.PublisherConfig{
		PriceDiffThreshold: 1,
		UpdateThreshold:    5 * time.Minute,
	}

	// Create publisher and call PublishUpdate
	publisher := NewUpdatePublisher(conf, []transactor.ITransactor{transactorMock1, transactorMock2}, fetcherMock, []string{dataKey}, []config.AssetSet{})
	err = publisher.PublishUpdate(ctx)
	r.NoError(err)

	assert.True(t, transactorMock1.AssertExpectations(t))
	assert.True(t, transactorMock2.AssertExpectations(t))
}

func TestNewMerkleUpdateFromProof(t *testing.T) {
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	assert.Nil(t, err)
	feedProofs.MerkleRoot = "0xd6e3f5da723db2bad4081e40f190076d8c579a0eac8e90703b06278b82a5e8f7"
	expectedMerkleRoot := []byte{0xd6, 0xe3, 0xf5, 0xda, 0x72, 0x3d, 0xb2, 0xba, 0xd4, 0x08, 0x1e, 0x40, 0xf1, 0x90, 0x07, 0x6d, 0x8c, 0x57, 0x9a, 0x0e, 0xac, 0x8e, 0x90, 0x70, 0x3b, 0x06, 0x27, 0x8b, 0x82, 0xa5, 0xe8, 0xf7}

	merkleUpdate, err := NewMerkleUpdateFromProof(&feedProofs)
	assert.Nil(t, err)
	recoveredDataKey := string(bytes.TrimRight(merkleUpdate.DataKey[:], "\x00"))
	assert.Equal(t, feedProofs.Key, recoveredDataKey)
	assert.Equal(t, expectedMerkleRoot, merkleUpdate.NewMerkleRoot[:])
	assert.Equal(t, len(feedProofs.MerkleProofs), len(merkleUpdate.MerkleProof))
	for i, proof := range feedProofs.MerkleProofs {
		var proofBytes [32]byte
		copy(proofBytes[:], proof)
		// assert.True(t, proof, merkleUpdate.MerkleProof[i][:])
		assert.Equal(t, proofBytes[:], merkleUpdate.MerkleProof[i][:])
	}

	assert.Equal(t, len(feedProofs.Signatures), len(merkleUpdate.Signatures))
	for i, sig := range feedProofs.Signatures {
		// assert.
		assert.Equal(t, sig.V, merkleUpdate.Signatures[i].V)
		assert.Equal(t, sig.R, merkleUpdate.Signatures[i].R)
		assert.Equal(t, sig.S, merkleUpdate.Signatures[i].S)
	}

	assert.Equal(t, bytes.TrimRight(feedProofs.Value.PriceData, "\x00"), merkleUpdate.Price.Bytes())
	assert.Equal(t, feedProofs.Value.Timestamp, merkleUpdate.Timestamp.Int64())
}

func TestUpdateTransactor_ValidateUpdate(t *testing.T) {
	r := require.New(t)

	key, err := crypto.GenerateKey()
	r.NoError(err)

	dataKey := "NGL/USDT"
	merkleUpdate, err := update.GenerateUpdate(key, dataKey)
	r.NoError(err)

	latestUpdate := merkleUpdate.Timestamp

	transactorMock := mockTransactor.NewMockITransactor(t)
	transactorMock.On("LatestUpdate", merkleUpdate.DataKey).Return(big.NewInt(200), latestUpdate)

	conf := config.PublisherConfig{
		PriceDiffThreshold: 500, // 5%
		UpdateThreshold:    5 * time.Minute,
	}

	updTransactor := newUpdateTransactor(conf, transactorMock)

	testCases := []struct {
		desc        string
		update      *types.MerkleRootUpdate
		containsErr string
		wantErr     bool
	}{
		{
			desc:    "Valid",
			update:  merkleUpdate, // with price diff 100% > threshold 5%
			wantErr: false,
		},
		{
			desc: "Too Often",
			update: &types.MerkleRootUpdate{
				DataKey:   merkleUpdate.DataKey,
				Timestamp: new(big.Int).Add(latestUpdate, big.NewInt(int64(conf.UpdateThreshold.Seconds()-1))),
				Price:     big.NewInt(200), // same price
			},
			containsErr: "too often update",
			wantErr:     true,
		},
		{
			desc: "Small Diff But Reach Threshold",
			update: &types.MerkleRootUpdate{
				DataKey:   merkleUpdate.DataKey,
				Timestamp: new(big.Int).Add(latestUpdate, big.NewInt(int64(conf.UpdateThreshold.Seconds()))),
				Price:     big.NewInt(210), // small changes ~4.76% < threshold 5%
			},
			wantErr: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := updTransactor.validateUpdate(tC.update)
			if tC.wantErr {
				r.Error(err)
				r.ErrorContains(err, tC.containsErr)
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestUpdateTransactor_SendUpdate(t *testing.T) {
	r := require.New(t)

	key, err := crypto.GenerateKey()
	r.NoError(err)

	dataKey := "NGL/USDT"
	merkleUpdate, err := update.GenerateUpdate(key, dataKey)
	r.NoError(err)

	transactorMock := mockTransactor.NewMockITransactor(t)
	transactorMock.On("SendUpdate", merkleUpdate).Return(nil, nil)
	transactorMock.On("LatestUpdate", merkleUpdate.DataKey).Return(big.NewInt(0), big.NewInt(0))

	conf := config.PublisherConfig{
		PriceDiffThreshold: 1,
		UpdateThreshold:    5 * time.Minute,
	}

	updTransactor := newUpdateTransactor(conf, transactorMock)

	_, ok := updTransactor.dataKeys[merkleUpdate.DataKey]
	r.False(ok)

	err = updTransactor.sendUpdate(merkleUpdate)
	r.NoError(err)

	latest, ok := updTransactor.dataKeys[merkleUpdate.DataKey]
	r.True(ok)
	r.Equal(latest.price, merkleUpdate.Price)
	r.Equal(latest.updatedAt, merkleUpdate.Timestamp)
}
