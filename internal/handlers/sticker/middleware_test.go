package stickerhandle

import (
	"testing"
	"tg-bot-sticker-go/internal/services/sticker"

	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)

func TestHandleOnText_HappyPath(t *testing.T){
	var msg telebot.Message
	msg.Text = "sticker"
	err := HandleOnText(&msg)
	assert.NoError(t, err)
}

func TestHandleOnText_FailCases(t *testing.T){
	test := []struct{
		Name string
		Message string
		ErrName error
	}{
		{
			Name: "empty Message",
			Message: "",
            ErrName: sticker.ErrInvalidPayload,
		},
		{
			Name: "invalid Message",
			Message: "INVALID-SET-NAME",
            ErrName: sticker.ErrInvalidPayload,
		},
	}
	for _, test := range test{
		t.Run(test.Name, func(t *testing.T){
            
            err := HandleOnText(&telebot.Message{
				Text: test.Message,
			})
            assert.Equal(t, test.ErrName, err)
        })
    }
}