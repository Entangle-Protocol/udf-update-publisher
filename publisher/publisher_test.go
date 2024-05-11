package publisher

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/brianvoe/gofakeit/v7"

	"gitlab.ent-dx.com/entangle/pull-update-publisher/mocks"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
)

func TestPublishUpdate(t *testing.T) {
	fetcherMock := mocks.NewMockIFetcher(t)

	// Generate test proofs and set mocker rv to return them
	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	fetcherMock.On("GetFeedProofs").Return(&feedProofs, nil)

	// Create publisher and call PublishUpdate
	publisher := NewUpdatePublisher(nil, fetcherMock)
	err = publisher.PublishUpdate()
	assert.Nil(t, err)
}
