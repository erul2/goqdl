package goqdl

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	baseURL        = "https://play.qobuz.com"
	loginURL       = baseURL + "/login"
	bundleRegexStr = `<script src="(/resources/\d+\.\d+\.\d+-[a-z]\d{3}/bundle\.js)"></script>`
	appIDRegexStr  = `production:{api:{appId:"(\d{9})",appSecret:"\w{32}"`
)

var (
	bundleRegex = regexp.MustCompile(bundleRegexStr)
	appIDRegex  = regexp.MustCompile(appIDRegexStr)
)

type bundle struct {
	client *http.Client
	bundle string
}

type Secrets map[string]string

func newBundle() (*bundle, error) {
	client := &http.Client{}
	resp, err := client.Get(loginURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch login page. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bundleMatch := bundleRegex.FindStringSubmatch(string(body))
	if len(bundleMatch) < 2 {
		return nil, fmt.Errorf("bundle URL not found")
	}

	bundleURL := baseURL + bundleMatch[1]

	resp, err = client.Get(bundleURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch bundle. Status code: %d", resp.StatusCode)
	}

	bundleBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &bundle{
		client: client,
		bundle: string(bundleBody),
	}, nil
}

func (b *bundle) GetAppID() (string, error) {
	appIDMatch := appIDRegex.FindStringSubmatch(b.bundle)
	if len(appIDMatch) < 2 {
		return "", fmt.Errorf("Failed to match APP ID")
	}

	return appIDMatch[1], nil
}

func (b *bundle) GetSecrets() (Secrets, error) {
	tzSecrets := make(map[string]string)
	seedTimezoneRegexStr := `[a-z]\.initialSeed\("([\w=]+)",window\.utimezone\.([a-z]+)\)`
	seedTimezoneRegex := regexp.MustCompile(seedTimezoneRegexStr)
	seedMatches := seedTimezoneRegex.FindAllStringSubmatch(b.bundle, -1)

	for _, match := range seedMatches {
		seed, timezone := match[1], match[2]
		tzSecrets[timezone] = seed
	}

	infoExtrasRegexStr := `name:"\w+/(?P<timezone>{timezones})",info:"([\w=]+)",extras:"([\w=]+)"`
	infoExtrasRegexStr = strings.Replace(infoExtrasRegexStr, "{timezones}", strings.Join(getTimezones(tzSecrets), "|"), -1)
	infoExtrasRegex := regexp.MustCompile(infoExtrasRegexStr)
	infoExtrasMatches := infoExtrasRegex.FindAllStringSubmatch(b.bundle, -1)

	secrets := make(Secrets, len(infoExtrasMatches))
	for _, match := range infoExtrasMatches {
		timezone, info, extras := match[1], match[2], match[3]
		secrets[strings.ToLower(timezone)] = decodeSecret(tzSecrets[strings.ToLower(timezone)] + info + extras)
	}

	return secrets, nil
}

func getTimezones(secrets Secrets) []string {
	timezones := make([]string, 0, len(secrets))
	for timezone := range secrets {
		timezones = append(timezones, cases.Title(language.English).String(timezone))
	}
	return timezones
}

func decodeSecret(secret string) string {
	if len(secret) <= 44 {
		return ""
	}
	secret = secret[:len(secret)-44]
	decoded, _ := base64.StdEncoding.DecodeString(secret)
	return string(decoded)
}
