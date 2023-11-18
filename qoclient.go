package goqdl

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	ErrFreeAccountNotEligible = errors.New("Free accounts are not eligible to download tracks.")
	ErrorSecretNotFound       = errors.New("Can't find any valid app secret.")
)

const (
	defaultBaseUrl = "https://www.qobuz.com/api.json/0.2"
	userAgent      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.207.132.170 Safari/537.36"
)

type credential struct {
	secrets  Secrets
	appId    string
	email    string
	password string
}

type qoClient struct {
	httpClient    *http.Client
	cred          credential
	secret        string
	baseUrl       string
	userAuthToken string
	label         string
}

func newQoClient(httpClient *http.Client, cred credential) (*qoClient, error) {
	ctx := context.Background()
	client := &qoClient{httpClient: httpClient, cred: cred, baseUrl: defaultBaseUrl}
	if err := client.login(ctx); err != nil {
		return nil, err
	}
	if err := client.setupSecret(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *qoClient) login(ctx context.Context) error {
	resp, err := func() (LoginResponse, error) {
		target := c.baseUrl + "/user/login"
		param := url.Values{
			"email":    []string{c.cred.email},
			"password": []string{c.cred.password},
			"app_id":   []string{c.cred.appId},
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewBufferString(param.Encode()))
		if err != nil {
			return LoginResponse{}, err
		}

		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("X-App-Id", c.cred.appId)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LoginResponse{}, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LoginResponse{}, err
		}

		if res.StatusCode != http.StatusOK {
			return LoginResponse{}, fmt.Errorf("Failed to login. Status code: %d Body: %s", res.StatusCode, string(body))
		}

		var resp LoginResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			return LoginResponse{}, err
		}

		return resp, nil
	}()
	if err != nil {
		return err
	}

	if resp.User.Credential.Parameters.ShortLabel == "" {
		return ErrFreeAccountNotEligible
	}

	c.userAuthToken = resp.UserAuthToken
	c.label = resp.User.Credential.Parameters.ShortLabel

	return nil
}

const (
	testTrackId = "5966783"
	testFmtId   = "5"
)

func (c *qoClient) getTrackUrl(ctx context.Context, trackID, formatID string) (GetFileURLResponse, error) {
	target := c.baseUrl + "/track/getFileUrl"
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)

	rsig := fmt.Sprintf("trackgetFileUrlformat_id%sintentstreamtrack_id%s%s%s", formatID, trackID, currentTime, c.secret)
	rsigHashed := md5.Sum([]byte(rsig))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return GetFileURLResponse{}, err
	}

	q := req.URL.Query()
	q.Set("request_ts", currentTime)
	q.Set("request_sig", fmt.Sprintf("%x", rsigHashed))
	q.Set("format_id", formatID)
	q.Set("track_id", trackID)
	q.Set("intent", "stream")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-App-Id", c.cred.appId)
	req.Header.Set("X-User-Auth-Token", c.userAuthToken)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GetFileURLResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetFileURLResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return GetFileURLResponse{}, fmt.Errorf("failed to get file URL Status code: %d Body: %s", res.StatusCode, string(body))
	}

	response := GetFileURLResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return GetFileURLResponse{}, fmt.Errorf("failed to get file URL: %s status code: %d Body: %s", err, res.StatusCode, string(body))
	}

	return response, nil
}

func (c *qoClient) setupSecret(ctx context.Context) error {
	for _, secret := range c.cred.secrets {
		c.secret = secret
		if _, err := c.getTrackUrl(ctx, testTrackId, testFmtId); err != nil {
			continue
		}
		return nil
	}
	return ErrorSecretNotFound
}

func (c *qoClient) getTrackMeta(ctx context.Context, id string) (GetTrackMetaResponse, error) {
	target := c.baseUrl + "/track/get"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return GetTrackMetaResponse{}, err
	}

	q := req.URL.Query()
	q.Set("track_id", id)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-App-Id", c.cred.appId)
	req.Header.Set("X-User-Auth-Token", c.userAuthToken)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GetTrackMetaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetTrackMetaResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return GetTrackMetaResponse{}, fmt.Errorf("failed to get track meta Status code: %d Body: %s", res.StatusCode, string(body))
	}

	response := GetTrackMetaResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return GetTrackMetaResponse{}, fmt.Errorf("failed to get track meta: %s status code: %d Body: %s", err, res.StatusCode, string(body))
	}

	return response, nil
}

func (c *qoClient) getAlbumMeta(ctx context.Context, id string) (GetAlbumMetaResponse, error) {
	target := c.baseUrl + "/album/get"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return GetAlbumMetaResponse{}, err
	}

	q := req.URL.Query()
	q.Set("album_id", id)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-App-Id", c.cred.appId)
	req.Header.Set("X-User-Auth-Token", c.userAuthToken)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GetAlbumMetaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return GetAlbumMetaResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return GetAlbumMetaResponse{}, fmt.Errorf("failed to get album meta Status code: %d Body: %s", res.StatusCode, string(body))
	}

	response := GetAlbumMetaResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return GetAlbumMetaResponse{}, fmt.Errorf("failed to get album meta: %s status code: %d Body: %s", err, res.StatusCode, string(body))
	}

	return response, nil
}
