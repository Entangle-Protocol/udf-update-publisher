package mocks

import (
	"testing"

	mockFetcher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
	// mockPublisher "gitlab.ent-dx.com/entangle/pull-update-publisher/mocks/gitlab.ent-dx.com/entangle/pull-update-publisher/publisher"
)

func NewMockIFetcher(t *testing.T) *mockFetcher.MockIFetcher {
	// t.Cleanup = func() { }
	return mockFetcher.NewMockIFetcher(t)
}
