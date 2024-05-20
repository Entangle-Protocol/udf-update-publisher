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
