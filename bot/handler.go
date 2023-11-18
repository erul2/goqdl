package bot

import (
	"context"
	"github.com/erul2/goqdl"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	qdl *goqdl.GoQDL
}

func NewBotHandler(qdl *goqdl.GoQDL) *Handler {
	return &Handler{qdl: qdl}
}

func (h *Handler) DownloadAlbum(c tele.Context) error {
	cmd := c.Message().Text
	err := h.qdl.HandleURL(context.Background(), cmd)
	if err != nil {
		return c.Send(err.Error())
	}
	return c.Send("download completed")
}
