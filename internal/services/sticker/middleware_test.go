package sticker

import (
	"testing"
	"tg-bot-sticker-go/internal/models/user"

	"github.com/stretchr/testify/assert"
)

func Test_CheckUserIDandName(t *testing.T){
	user := user.User{
		ID: "123456",
        UserName: "username",
	}
	err := CheckUserIDandName(&user)
	assert.NoError(t, err)
}

func Test_CheckEmojiandSetName(t *testing.T){
	user := user.User{
		Emoji: "123456",
        SetName: "Setname",
	}
	err := CheckEmojiandSetName(&user)
	assert.NoError(t, err)
}

func Test_CheckSetName(t *testing.T){
	user := user.User{
        SetName: "Setname",
    }
    err := CheckSetName(&user)
    assert.NoError(t, err)
}

func Test_CheckUserIDandName_FailCases(t *testing.T){
	
	test := []struct{
		Name string
        ID string
		UserName string
		ErrName error
	}{
		{
			Name: "empty ID",
			ID: "",
            UserName: "user",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "empty UserName",
			ID: "123456",
            UserName: "",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "empty Both ID and UserName",
			ID: "",
            UserName: "",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "invalid ID",
			ID: "invalid-id",
            UserName: "Name",
            ErrName: ErrInvalidPayload,
		},
	}
	for _, test := range test{
		t.Run(test.Name, func(t *testing.T){
            user := user.User{
                ID: test.ID,
                UserName: test.UserName,
            }
            err := CheckUserIDandName(&user)
            assert.Equal(t, test.ErrName, err)
        })
    }
}

func Test_CheckEmojiandSetName_FailCases(t *testing.T){
	test := []struct{
		Name string
        Emoji string
		SetName string
		ErrName error
	}{
		{
			Name: "empty Emoji",
			Emoji: "",
            SetName: "user",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "empty SetName",
			Emoji: "123456",
            SetName: "",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "empty Both Emoji and SetName",
			Emoji: "",
            SetName: "",
            ErrName: ErrInvalidPayload,
		},
	}
	for _, test := range test{
		t.Run(test.Name, func(t *testing.T){
            user := user.User{
                Emoji: test.Emoji,
                SetName: test.SetName,
            }
            err := CheckUserIDandName(&user)
            assert.Equal(t, test.ErrName, err)
        })
    }
}

func Test_CheckSetName_FailCases(t *testing.T){
	test := []struct{
		Name string
		SetName string
		ErrName error
	}{
		{
			Name: "empty SetName",
			SetName: "",
            ErrName: ErrInvalidPayload,
		},
		{
			Name: "invalid UserName",
			SetName: "INVALID-SET-NAME",
            ErrName: ErrInvalidPayload,
		},
	}
	for _, test := range test{
		t.Run(test.Name, func(t *testing.T){
            user := user.User{
                SetName: test.SetName,
            }
            err := CheckUserIDandName(&user)
            assert.Equal(t, test.ErrName, err)
        })
    }
}