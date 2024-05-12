package publisher

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"

	mockFetcher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	mockTransactor "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/transactor"
	// mockPublisher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
)

func TestPublishUpdate(t *testing.T) {
	fetcherMock := mockFetcher.NewMockIFetcher(t)

	// Generate test proofs and set mocker rv to return them
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	fetcherMock.On("GetFeedProofs").Return(&feedProofs, nil)

	transactorMock := mockTransactor.NewMockITransactor(t)
	merkleUpdate := NewMerkleUpdateFromProof(&feedProofs)
	transactorMock.On("SendUpdate", merkleUpdate).Return(nil)

	// Create publisher and call PublishUpdate
	publisher := NewUpdatePublisher(transactorMock, fetcherMock)
	err = publisher.PublishUpdate()
	assert.Nil(t, err)

	assert.True(t, transactorMock.AssertExpectations(t))
}

func TestNewMerkleUpdateFromProof(t *testing.T) {
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	assert.Nil(t, err)

	merkleUpdate := NewMerkleUpdateFromProof(&feedProofs)
	assert.Equal(t, feedProofs.MerkleRoot.Bytes(), merkleUpdate.NewMerkleRoot[:])
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
