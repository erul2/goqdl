package goqdl

import (
	"context"
	"fmt"
	"net/http"
)

type GoQDLConfig struct {
	DownloadDir     string
	DownloadQuality string
}
type GoQDL struct {
	httpClient *http.Client
	qoClient   *qoClient
	config     GoQDLConfig
	creds      credential
}

func NewGoQDL(email string, password string, config GoQDLConfig) (*GoQDL, error) {
	qdl := &GoQDL{
		httpClient: &http.Client{},
		config:     config,
		creds: credential{
			email:    email,
			password: password,
		},
	}
	if err := qdl.initializeSecrets(); err != nil {
		return nil, err
	}
	if err := qdl.initializeClient(); err != nil {
		return nil, err
	}
	return qdl, nil
}

func (g *GoQDL) initializeClient() error {
	client, err := newQoClient(
		g.httpClient,
		g.creds,
	)
	if err != nil {
		return err
	}
	g.qoClient = client
	return nil
}

func (g *GoQDL) initializeSecrets() error {
	bundle, err := newBundle()
	if err != nil {
		return err
	}
	appID, err := bundle.GetAppID()
	if err != nil {
		return err
	}
	g.creds.appId = appID

	secrets, err := bundle.GetSecrets()
	if err != nil {
		return err
	}
	g.creds.secrets = secrets
	return nil
}

func (g *GoQDL) HandleURL(ctx context.Context, url string) error {
	itemType, itemID := getURLInfo(url)
	switch itemType {
	// case "album", "track":
	case "album":
		return g.downloadAlbum(ctx, itemID)
	case "playlist", "artist", "label":
	default:
		return fmt.Errorf("unknown item type: %s", itemType)
	}
	return nil
}

func (g *GoQDL) downloadAlbum(ctx context.Context, id string) error {
	downloader := newDownloader(5, g.qoClient, g.config.DownloadDir, id, g.config.DownloadQuality)
	return downloader.downloadRelease(context.Background())
}
