package sticker

import (
	"context"
	"runtime"
	"testing"
	"tg-bot-sticker-go/internal/services/sticker/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreate_HappyPath(t *testing.T) {
	runtime.Gosched()
	chatID := "12345"
	user := []byte("12345 12345 1 1")
	rep := mocks.NewRepository(t)
	
	rep.On("Create",  context.Background(),chatID, user).Return(nil).Once()
	
	err := rep.Create(context.Background(), chatID, user)
	assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T) {
	runtime.Gosched()
	chatID := "12345"
	rep := mocks.NewRepository(t)
	var userBytes []byte
	var err error
	
	rep.On("Get",  context.Background(),chatID).Return(userBytes, nil).Once()
	
	_, err = rep.Get(context.Background(), chatID)
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T) {
	runtime.Gosched()
	chatID := "12345"
	user := []byte("12345 12345 1 1")
	rep := mocks.NewRepository(t)
	
	rep.On("Update",  context.Background(),chatID, user).Return(nil).Once()
	
	err := rep.Update(context.Background(), chatID, user)
	assert.NoError(t, err)
}

func TestIsExist_HappyPath(t *testing.T) {
	runtime.Gosched()
	chatID := "12345"
	var isExist int64

	rep := mocks.NewRepository(t)
	
	rep.On("IsExist",  context.Background(),chatID).Return(isExist, nil).Once()
	
	_, err := rep.IsExist(context.Background(), chatID)
	assert.NoError(t, err)
}

func TestDelete_HappyPath(t *testing.T) {
	runtime.Gosched()
	chatID := "12345"
	
	rep := mocks.NewRepository(t)
	
	rep.On("Delete",  context.Background(),chatID).Return(nil).Once()
	
	err := rep.Delete(context.Background(), chatID)
	assert.NoError(t, err)
}

func TestCreate_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestGet_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestUpdate_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestIsExist_FailCases(t *testing.T) {
	// TODO: implement me.
}

func TestDelete_FailCases(t *testing.T) {
	// TODO: implement me.
}