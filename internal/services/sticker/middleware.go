package sticker

import (
	"strings"
	"tg-bot-sticker-go/internal/models/user"
)

const (
	invalidSymbolsID = "=_!@#$%^&*()-=+/?.,`\\\"<>№;:[]{|} abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	invalidSymbolsSetName = "=_!@#$%^&*()-=+/?.,`\\\"<>№;:[]{|} "
)

func CheckUserIDandName(user *user.User) error{
	if user.ID == "" || user.UserName == "" {
        return ErrInvalidPayload
	}
	if strings.ContainsAny(user.ID, invalidSymbolsID){
		return ErrInvalidPayload
	}
	return nil
}

func CheckEmojiandSetName(user *user.User) error{
	if user.Emoji == "" || user.SetName == ""{
		return ErrInvalidPayload
	}
	if strings.ContainsAny(user.SetName, invalidSymbolsSetName){
		return ErrInvalidPayload
	}
	return nil
}

func CheckSetName(user *user.User) error{
	if user.SetName == "" {
		return ErrInvalidPayload
	}
	if strings.ContainsAny(user.SetName, invalidSymbolsSetName){
		return ErrInvalidPayload
	}
	return nil
}