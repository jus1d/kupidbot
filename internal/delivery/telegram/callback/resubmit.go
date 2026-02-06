package callback

import (
	"context"
	"log/slog"

	"github.com/jus1d/kypidbot/internal/delivery/telegram/view"
	"github.com/jus1d/kypidbot/internal/lib/logger/sl"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) Resubmit(c tele.Context) error {
	sender := c.Sender()

	if err := h.Registration.SetState(context.Background(), sender.ID, "awaiting_sex"); err != nil {
		slog.Error("set state", sl.Err(err))
		return c.Respond()
	}

	return h.DeleteAndSend(c, view.Msg("start", "ask_sex", "resubmit"), view.SexKeyboard())
}
