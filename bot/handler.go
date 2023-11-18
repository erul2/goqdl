package bot

import (
	"context"
	"fmt"
	"github.com/erul2/goqdl"
	tele "gopkg.in/telebot.v3"
	"regexp"
)

type Handler struct {
	qdl *goqdl.GoQDL
}

func NewBotHandler(qdl *goqdl.GoQDL) *Handler {
	return &Handler{qdl: qdl}
}

func extractURLFromString(input string) (string, error) {
	// Regular expression pattern to match URLs
	urlPattern := `https?://[^\s]+`

	// Compile the regular expression pattern
	r, err := regexp.Compile(urlPattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile regular expression: %w", err)
	}

	// Find the first match in the input string
	match := r.FindString(input)
	if match == "" {
		return "", fmt.Errorf("no URL found in the input string")
	}

	return match, nil
}

func (h *Handler) DownloadAlbum(c tele.Context) error {
	cmd := c.Message().Text
	url, err := extractURLFromString(cmd)
	if err != nil {
		return c.Send("url not found")
	}

	err = h.qdl.HandleURL(context.Background(), url)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Send("download completed")
}
