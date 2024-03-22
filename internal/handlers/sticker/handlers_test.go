package stickerhandle

import (
	"context"
	"testing"
	"tg-bot-sticker-go/internal/handlers/sticker/mocks"
	"tg-bot-sticker-go/internal/models/actions"

	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)



func TestPrepare_HappyPath(t *testing.T) {
	st := mocks.NewSticker(t)

	st.On("Prepare",  context.Background(), &telebot.Message{}, &telebot.Bot{}, 1).Return(nil).Once()

	err := st.Prepare(context.Background(), &telebot.Message{}, &telebot.Bot{}, 1)
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T) {
	st := mocks.NewSticker(t)

	st.On("Update",  context.Background(), &telebot.Message{}, &telebot.Bot{}).Return(actions.Actions{}, nil).Once()

	_, err := st.Update(context.Background(), &telebot.Message{}, &telebot.Bot{})
	assert.NoError(t, err)
}

func TestDone_HappyPath(t *testing.T) {
	st := mocks.NewSticker(t)

	st.On("Done",  context.Background(), &telebot.Message{}, &telebot.Bot{}).Return(0, nil).Once()

	_, err := st.Done(context.Background(), &telebot.Message{}, &telebot.Bot{})
	assert.NoError(t, err)
}

func TestPrepare_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestUpdate_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestDone_FailCases(t *testing.T) {
	// TODO: implement me.
}