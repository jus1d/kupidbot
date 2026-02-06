package callback

import (
	"context"
	"log/slog"

	"github.com/jus1d/kypidbot/internal/config/messages"
	"github.com/jus1d/kypidbot/internal/lib/logger/sl"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) Sex(c tele.Context) error {
	sender := c.Sender()
	cb := c.Callback()

	sex := "female"
	if cb.Unique == "sex_male" {
		sex = "male"
	}

	if err := h.Registration.SetSex(context.Background(), sender.ID, sex); err != nil {
		slog.Error("set sex", sl.Err(err))
		return c.Respond()
	}

	if err := h.Registration.SetState(context.Background(), sender.ID, "awaiting_about"); err != nil {
		slog.Error("set state", sl.Err(err))
		return c.Respond()
	}

	return h.DeleteAndSend(c, messages.M.Profile.About.Request)
}
