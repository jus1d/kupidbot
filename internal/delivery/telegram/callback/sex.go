package callback

import (
	"context"

	"github.com/jus1d/kypidbot/internal/delivery/telegram/view"
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
		h.Log.Error("set sex", sl.Err(err))
		return c.Respond()
	}

	if err := h.Registration.SetState(context.Background(), sender.ID, "awaiting_about"); err != nil {
		h.Log.Error("set state", sl.Err(err))
		return c.Respond()
	}

	return c.Edit(view.Msg("sex_selected"))
}
