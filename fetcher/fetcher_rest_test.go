package fetcher_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"gitlab.ent-dx.com/entangle/pull-update-publisher/fetcher"
)

func Test_GetFeedProofs_Rest(t *testing.T) {
	r := require.New(t)

	var feedProofs fetcher.EntangleFeedProof
	err := gofakeit.Struct(&feedProofs)
	r.NoError(err)

	testCases := []struct {
		desc     string
		handler  http.HandlerFunc
		assetKey string
		wantErr  bool
	}{
		{
			desc:     "Valid",
			assetKey: "NGL/USDT",
			handler: func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				m := map[string]any{
					"calldata": feedProofs,
					"error":    "",
				}
				json.NewEncoder(w).Encode(&m)
			},
			wantErr: false,
		},
		{
			desc: "Error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				m := map[string]any{
					"error": "failed to retrieve data feed proof",
				}

				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(&m)
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			srv := httptest.NewServer(tC.handler)
			fetcher := fetcher.NewRestFetcher(http.DefaultClient, srv.URL)

			gotProof, err := fetcher.GetFeedProofs(context.Background(), tC.assetKey)

			srv.Close()

			if tC.wantErr {
				r.Error(err)
			} else {
				r.NotEmpty(gotProof)

				r.Equal(feedProofs.Key, gotProof.Key)
				r.Equal(feedProofs.Value, gotProof.Value)
				r.Equal(feedProofs.MerkleRoot, gotProof.MerkleRoot)
				r.Equal(feedProofs.MerkleProofs, gotProof.MerkleProofs)
				r.Equal(feedProofs.Signatures, gotProof.Signatures)
			}
		})
	}
}
func Test_GetSpotterFeedsProofs_Rest(t *testing.T) {
	r := require.New(t)

	var feedsProofs fetcher.EntangleFeedsProofs
	err := gofakeit.Struct(&feedsProofs)
	r.NoError(err)

	testCases := []struct {
		desc      string
		handler   http.HandlerFunc
		spotterID string
		assetKeys []string
		wantErr   bool
	}{
		{
			desc:      "Valid",
			spotterID: "spotter1",
			assetKeys: []string{"NGL/USDT", "BTC/USD"},
			handler: func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				m := map[string]interface{}{
					"calldata": feedsProofs,
					"error":    "",
				}
				json.NewEncoder(w).Encode(&m)
			},
			wantErr: false,
		},
		{
			desc:      "Error",
			spotterID: "spotter2",
			assetKeys: []string{"ETH/USD", "BTC/USD"},
			handler: func(w http.ResponseWriter, r *http.Request) {
				m := map[string]interface{}{
					"error": "failed to retrieve spotter feeds proofs",
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(&m)
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			srv := httptest.NewServer(tC.handler)
			fetcher := fetcher.NewRestFetcher(http.DefaultClient, srv.URL)

			gotProofs, err := fetcher.GetSpotterFeedsProofs(context.Background(), tC.spotterID, tC.assetKeys)
			srv.Close()

			if tC.wantErr {
				r.Error(err)
			} else {
				r.NotEmpty(gotProofs)
				r.Equal(feedsProofs.MerkleRoot, gotProofs.MerkleRoot)
				r.EqualValues(feedsProofs.Signatures, gotProofs.Signatures)
				r.EqualValues(feedsProofs.Feeds, gotProofs.Feeds)
			}
		})
	}
}
