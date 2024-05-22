package publisher

import (
	"bytes"
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"

	mockFetcher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	mockTransactor "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	// mockPublisher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
)

func TestPublishUpdate(t *testing.T) {
	ctx := context.Background()
	fetcherMock := mockFetcher.NewMockIFetcher(t)

	// Generate test proofs and set mocker rv to return them
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	assert.Nil(t, err)

	dataKey := "NGL/USDT"
	fetcherMock.On("GetFeedProofs", ctx, dataKey).Return(&feedProofs, nil)

	transactorMock1 := mockTransactor.NewMockITransactor(t)
	transactorMock2 := mockTransactor.NewMockITransactor(t)

	merkleUpdate, err := NewMerkleUpdateFromProof(&feedProofs)
	assert.Nil(t, err)
	transactorMock1.On("SendUpdate", merkleUpdate).Return(nil)
	transactorMock2.On("SendUpdate", merkleUpdate).Return(nil)

	// Create publisher and call PublishUpdate
	publisher := NewUpdatePublisher([]transactor.ITransactor{transactorMock1, transactorMock2}, fetcherMock, []string{dataKey})
	err = publisher.PublishUpdate(ctx)
	assert.Nil(t, err)

	assert.True(t, transactorMock1.AssertExpectations(t))
	assert.True(t, transactorMock2.AssertExpectations(t))
}

func TestNewMerkleUpdateFromProof(t *testing.T) {
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	assert.Nil(t, err)

	merkleUpdate, err := NewMerkleUpdateFromProof(&feedProofs)
	assert.Nil(t, err)
	recoveredDataKey := string(bytes.TrimRight(merkleUpdate.DataKey[:], "\x00"))
	assert.Equal(t, feedProofs.Key, recoveredDataKey)
	assert.Equal(t, feedProofs.MerkleRoot, merkleUpdate.NewMerkleRoot[:])
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

	assert.Equal(t, feedProofs.Value.PriceData, merkleUpdate.Price.Bytes())
	assert.Equal(t, feedProofs.Value.Timestamp, merkleUpdate.Timestamp.Int64())
}
