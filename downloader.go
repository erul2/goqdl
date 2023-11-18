package goqdl

import (
	"context"
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"os"
	"strconv"
)

type downloader struct {
	numWorker  int
	client     *qoClient
	grabClient *grab.Client
	rootDir    string
	itemID     string
	quality    string
}

func newDownloader(numWorker int, client *qoClient, rootDir string, itemID string, quality string) *downloader {
	grabClient := grab.NewClient()
	return &downloader{numWorker: numWorker, client: client, rootDir: rootDir, grabClient: grabClient, itemID: itemID, quality: quality}
}

type downloadRequest struct {
	url    string
	target string
}

func (d *downloader) download(requests ...downloadRequest) error {
	respCh := d.grabClient.DoBatch(d.numWorker, func() []*grab.Request {
		reqs := make([]*grab.Request, len(requests))
		for i := range requests {
			reqs[i], _ = grab.NewRequest(requests[i].target, requests[i].url)
		}
		return reqs
	}()...)
	for resp := range respCh {
		if err := resp.Err(); err != nil {
			return err
		}
	}
	return nil
}

func createFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %w", err)
		}
	}

	return nil
}

func (d *downloader) downloadRelease(ctx context.Context) error {
	meta, err := d.client.getAlbumMeta(ctx, d.itemID)
	if err != nil {
		return err
	}

	sampleTrack, err := d.client.getTrackUrl(ctx, strconv.Itoa(meta.Tracks.Items[0].Id), d.quality)
	if err != nil {
		return err
	}

	albumAttributes := meta.GetAlbumAttributes(sampleTrack)
	folder := fmt.Sprintf("%s/%s", d.rootDir, albumAttributes.folderFormat())
	if err := createFolderIfNotExists(folder); err != nil {
		return err
	}

	requests := func() []downloadRequest {
		reqs := make([]downloadRequest, len(meta.Tracks.Items))
		for i := range meta.Tracks.Items {
			track, err := d.client.getTrackUrl(ctx, strconv.Itoa(meta.Tracks.Items[i].Id), d.quality)
			if err != nil {
				return nil
			}
			fileName := fmt.Sprintf("%02d. %s.flac", i+1, meta.Tracks.Items[i].Title)
			reqs[i] = downloadRequest{
				url:    track.Url,
				target: folder + "/" + fileName,
			}
		}
		return reqs
	}()

	if err := d.download(requests...); err != nil {
		return err
	}

	return nil
}
