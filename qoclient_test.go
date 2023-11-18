package goqdl

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQoClientNewClient(t *testing.T) {
	qoClient, err := newQoClient(
		&http.Client{},
		credential{
			secrets: Secrets{
				"berlin": "979549437fcc4a3faad4867b5cd25dcb",
				"london": "10b251c286cfbf64d6b7105f253d9a2e",
			},
			appId:    "950096963",
			email:    "hioxlpye@uploadplaystore.com",
			password: "Pss9Oz199",
		},
	)
	assert.NoError(t, err)

	t.Run("getTrackURL", func(t *testing.T) {
		res, err := qoClient.getTrackUrl(context.Background(), "24176681", "7")
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("getTrackMeta", func(t *testing.T) {
		res, err := qoClient.getTrackMeta(context.Background(), "24176681")
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("getAlbumMeta", func(t *testing.T) {
		res, err := qoClient.getAlbumMeta(context.Background(), "0886971387827")
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
}
