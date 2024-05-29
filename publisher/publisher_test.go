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
	feedProofs.MerkleRoot = "0xd6e3f5da723db2bad4081e40f190076d8c579a0eac8e90703b06278b82a5e8f7"

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

	assert.Equal(t, feedProofs.Value.PriceData, merkleUpdate.Price.Bytes())
	assert.Equal(t, feedProofs.Value.Timestamp, merkleUpdate.Timestamp.Int64())
}
