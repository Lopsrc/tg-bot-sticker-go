package stickerhandle

import (
	"context"
	"errors"
	"log/slog"

	"tg-bot-sticker-go/internal/models/actions"
	"tg-bot-sticker-go/internal/services/sticker"
	"tg-bot-sticker-go/internal/utils/info"

	tele "gopkg.in/telebot.v3"
)

// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Sticker
type Sticker interface {
	// Prepare user.
	Prepare(ctx context.Context, msg *tele.Message, b *tele.Bot, code int) error
	// Update user information.
	Update(ctx context.Context, msg *tele.Message, b *tele.Bot) (code actions.Actions, err error)
	// Done action.
	Done(ctx context.Context, msg *tele.Message, b *tele.Bot) (int, error)
}
var(
	btnHelp = tele.Btn{
		Text: "help",
		Data: "Data",
	}
	btnCreateSet = tele.Btn{
		Text: "Create sticker set",
		Data: "Create",
	}
	btnAddSticker = tele.Btn{
		Text: "Add sticker to set",
		Data: "Add",
	}
	btnGetSticker = tele.Btn{
        Text: "Get sticker set",
        Data: "Get",
    }
	btnDeleteSicker = tele.Btn{
		Text: "Delete sticker",
        Data: "Delete",
    }
	btnDone = tele.Btn{
		Text: "Done",
		Data: "Data",
	}	
	btnCancel = tele.Btn{
		Text: "Cancel",
        Data: "Data",
    }
)

type Handler struct {
	s   Sticker
	log *slog.Logger
	bot *tele.Bot
}

// New creates a new instance of the handler.
func New(bot *tele.Bot, r sticker.Repository, log *slog.Logger) *Handler {

	s := sticker.New(r, log)
	return &Handler{
		s:   s,
		bot: bot,
		log: log,
	}
}

// Register registers all the handlers for the bot.
func (h *Handler) Register() {
	h.bot.Handle("/start", h.Start())
	h.bot.Handle(&btnHelp, h.Help())
	h.bot.Handle(&btnCreateSet, h.StickerSet())
	h.bot.Handle(&btnAddSticker, h.Add())
	h.bot.Handle(&btnGetSticker, h.Get())
	h.bot.Handle(&btnDeleteSicker, h.Delete())
	h.bot.Handle(&btnDone, h.Done())
	h.bot.Handle(&btnCancel, h.GetButtons())

	h.bot.Handle(tele.OnDocument, h.Done())
	h.bot.Handle(tele.OnPhoto, h.Done())
	h.bot.Handle(tele.OnText, h.Update())
	h.bot.Handle(tele.OnSticker, h.Done())
	
}

// Start sends a message with the start text to the user.
func (h *Handler) Start() tele.HandlerFunc {
	return func(ctx tele.Context) error {

		r := h.bot.NewMarkup()
		r.ResizeKeyboard = true
		r.Reply(
			r.Row(btnCreateSet),
			r.Row(btnAddSticker),
            r.Row(btnGetSticker),
            r.Row(btnDeleteSicker),
			r.Row(btnHelp),
		)

		h.bot.Send(ctx.Message().Sender, info.MsgStart, r)
		return nil
	}
}

// Help sends a message with the help text to the user.
func (h *Handler) Help() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		h.bot.Send(ctx.Message().Sender, info.MsgHelp)
		return nil
	}
}

// StickerSet prepares the user for creating a sticker set.
func (h *Handler) StickerSet() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		if err := h.s.Prepare(context.TODO(), ctx.Message(), h.bot, sticker.CreateStickerSet); err != nil {
			h.bot.Send(ctx.Message().Sender, info.MsgInternalError)
			return nil
		}
		h.bot.Send(ctx.Message().Sender, info.MsgCreateStickerSet)
		return nil
	}
}

// Add prepares the user for adding a sticker to the set.
func (h *Handler) Add() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		if err := h.s.Prepare(context.TODO(), ctx.Message(), h.bot, sticker.AddSticker); err != nil {
			h.bot.Send(ctx.Message().Sender, info.MsgInternalError)
			return nil
		}
		h.bot.Send(ctx.Message().Sender, info.MsgAddSticker)
		return nil
	}
}

// Get prepares the user for getting the sticker set.
func (h *Handler) Get() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		if err := h.s.Prepare(context.TODO(), ctx.Message(), h.bot, sticker.GetStickerSet); err != nil {
			h.bot.Send(ctx.Message().Sender, info.MsgInternalError)
			return nil
		}
		h.bot.Send(ctx.Message().Sender, info.MsgDeleteOrGet)
		return nil
	}
}

// Delete prepares the user for deleting a sticker from the set.
func (h *Handler) Delete() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		if err := h.s.Prepare(context.TODO(), ctx.Message(), h.bot, sticker.DeleteSticker); err != nil {
			h.bot.Send(ctx.Message().Sender, info.MsgCreateStickerSet)
			return nil
		}
		h.bot.Send(ctx.Message().Sender, info.MsgDeleteOrGet)
		return nil
	}
}

// Update updates the entity in the repository.
func (h *Handler) Update() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		// Handle a text.
		if err := HandleOnText(ctx.Message()); err != nil {
			h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
            return nil
        }
		// Update entity to the redis.
		codeAction, err := h.s.Update(context.TODO(), ctx.Message(), h.bot)
		if err != nil {
			if errors.Is(err, sticker.ErrInvalidPayload) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			} else if errors.Is(err, sticker.ErrEntityNotFound) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			}
			h.bot.Send(ctx.Message().Sender, info.MsgInternalError)
			return nil
		}
		// Detection of the action.
		switch codeAction.Code {
		case sticker.CreateStickerSet:
			if codeAction.Iter == 1{
				h.bot.Send(ctx.Message().Sender, info.MsgEmoji)
			}else {
				h.bot.Send(ctx.Message().Sender, info.MsgUploadPhoto)
			}
		case sticker.AddSticker:
			if codeAction.Iter == 1{
				h.bot.Send(ctx.Message().Sender, info.MsgEmoji)
			}else if codeAction.Iter >= 2 {
				h.bot.Send(ctx.Message().Sender, info.MsgUploadPhoto)
			}
		case sticker.GetStickerSet:
			r := h.bot.NewMarkup()
			r.ResizeKeyboard = true
			r.Reply(
				r.Row(btnDone),
				r.Row(btnCancel),
			)

			h.bot.Send(ctx.Message().Sender, info.MsgUpGet, r)
		case sticker.DeleteSticker:
			h.bot.Send(ctx.Message().Sender, info.MsgDeleteSticker)
		}
		//Reply the information.
		return nil
	}
}

// Done prepares the user for getting the sticker set.
func (h *Handler) Done() tele.HandlerFunc {
	return func(ctx tele.Context) error {
		codeAction, err := h.s.Done(context.TODO(), ctx.Message(), h.bot)
		if err != nil {
			if errors.Is(err, sticker.ErrInvalidPayload) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			} else if errors.Is(err, sticker.ErrStickersIsDeleted) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrStickersNotFound)
				return nil
			} else if errors.Is(err, sticker.ErrEntityNotFound) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			} else if errors.Is(err, tele.ErrStickerEmojisInvalid) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			} else if errors.Is(err, tele.ErrStickerSetInvalidName) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrPayload)
				return nil
			} else if errors.Is(err, tele.ErrStickerSetNameOccupied) {
				h.bot.Send(ctx.Message().Sender, info.MsgErrStickerSetExist)
				return nil
			}
			h.bot.Send(ctx.Message().Sender, info.MsgInternalError)
			return nil
		}
		// Detection of the action.
		switch codeAction {
		case sticker.CreateStickerSet:
			h.bot.Send(ctx.Message().Sender, info.MsgDoneCreate)
		case sticker.AddSticker:
			h.bot.Send(ctx.Message().Sender, info.MsgDoneAdd)
		case sticker.GetStickerSet:
			r := h.bot.NewMarkup()
			r.ResizeKeyboard = true
			r.Reply(
				r.Row(btnCreateSet),
				r.Row(btnAddSticker),
				r.Row(btnGetSticker),
				r.Row(btnDeleteSicker),
				r.Row(btnHelp),
			)
			h.bot.Send(ctx.Message().Sender, info.MsgDoneGet, r)
		case sticker.DeleteSticker:
			h.bot.Send(ctx.Message().Sender, info.MsgDoneDelete)
		}
		return nil
	}
}

func (h *Handler) GetButtons() tele.HandlerFunc{
	return func(ctx tele.Context) error {
        r := h.bot.NewMarkup()
        r.ResizeKeyboard = true
        r.Reply(
            r.Row(btnCreateSet),
            r.Row(btnAddSticker),
            r.Row(btnGetSticker),
            r.Row(btnDeleteSicker),
            r.Row(btnHelp),
        )
        h.bot.Send(ctx.Message().Sender, info.MsgHelp, r)
        return nil
    }
}