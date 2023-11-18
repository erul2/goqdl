package goqdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBundle(t *testing.T) {
	bundel, err := newBundle()
	if err != nil {
		t.Fatalf("TestBundle - NewBundle: %v", err)
	}

	t.Run("TestBundle_GetAppID", func(t *testing.T) {
		appId, err := bundel.GetAppID()
		assert.NoError(t, err)
		assert.NotEmpty(t, appId)
		assert.Equal(t, "950096963", appId)
	})

	t.Run("TestBundle_GetSecrets", func(t *testing.T) {
		secrets, err := bundel.GetSecrets()
		assert.NoError(t, err)
		assert.NotEmpty(t, secrets)
		assert.Equal(t, Secrets{
			"berlin": "979549437fcc4a3faad4867b5cd25dcb",
			"london": "10b251c286cfbf64d6b7105f253d9a2e",
		}, secrets)
	})
}
