package bot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractURLFromString(t *testing.T) {
	str := "Listen to the release LIKE IT LIKE IT by SECRET NUMBER on Qobuz https://open.qobuz.com/album/b71isise1ctta"
	url, err := extractURLFromString(str)
	assert.NoError(t, err)
	assert.Equal(t, "https://open.qobuz.com/album/b71isise1ctta", url)
}
