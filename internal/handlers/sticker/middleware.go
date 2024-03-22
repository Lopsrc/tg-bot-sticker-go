package stickerhandle

import (
	"strings"
	"tg-bot-sticker-go/internal/services/sticker"

	tele "gopkg.in/telebot.v3"
)

const (
	invalidSymbols = "=_!@#$%^&*()-=+/?.,<>`\"<>№;:[]{|} "
	maxPhotoSize = 64000 //64KB
)

func HandleOnText(msg *tele.Message) error{
	if strings.ContainsAny(msg.Text, invalidSymbols) {
        return sticker.ErrInvalidPayload
    }
	if msg.Text == "" {
        return sticker.ErrInvalidPayload
    }
	return nil
}
