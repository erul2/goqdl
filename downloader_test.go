package goqdl

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_downloader_downloadRelease(t *testing.T) {
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
	start := time.Now()
	downloader := newDownloader(4, qoClient, "./downloads", "mtwm9q5wt4pja", "5")
	err = downloader.downloadRelease(context.Background())
	log.Println("time elapsed ", time.Since(start).Seconds())
	assert.NoError(t, err)
}

// 2023/11/18 23:53:52 time elapsed  11.908669202
