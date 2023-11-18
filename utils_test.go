package goqdl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getURLInfo(t *testing.T) {
	tests := []struct {
		url          string
		expectedType string
		expectedID   string
	}{
		{
			url:          "https://open.qobuz.com/track/24176681",
			expectedType: "track",
			expectedID:   "24176681",
		},
		{
			url:          "https://open.qobuz.com/artist/578286",
			expectedType: "artist",
			expectedID:   "578286",
		},
		{
			url:          "https://open.qobuz.com/album/mtwm9q5wt4pja",
			expectedType: "album",
			expectedID:   "mtwm9q5wt4pja",
		},
		{
			url:          "https://open.qobuz.com/playlist/17925398",
			expectedType: "playlist",
			expectedID:   "17925398",
		},
	}
	for _, tt := range tests {
		actualType, actualID := getURLInfo(tt.url)
		assert.Equalf(t, tt.expectedType, actualType, "getURLInfo(%v)", tt.expectedType)
		assert.Equalf(t, tt.expectedID, actualID, "getURLInfo(%v)", tt.expectedID)

	}
}
