package sticker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"tg-bot-sticker-go/internal/models/actions"
	"tg-bot-sticker-go/internal/models/user"
	"tg-bot-sticker-go/internal/storage"
	"tg-bot-sticker-go/internal/utils/file"

	tele "gopkg.in/telebot.v3"
)


// go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Repository
type Repository interface {
	Create(ctx context.Context, chatID string, user []byte) error
	Get(ctx context.Context, chatID string) (user []byte, err error)
	Update(ctx context.Context, chatID string, user []byte) error
	IsExist(ctx context.Context, chatID string) (int64, error)
	Delete(ctx context.Context, chatID string) error
}

var (
	ErrEntityNotFound       = errors.New("entity not found")
	ErrInternalError        = errors.New("internal error")
	ErrInvalidPayload       = errors.New("invalid payload")
	ErrInvalidPhotoSize     = errors.New("invalid photo size")
	ErrInvalidFileExtension = errors.New("invalid file extension")
	ErrStickersIsDeleted	= errors.New("stickers is deleted")
)

const (
	CreateStickerSet = iota + 1
	AddSticker
	GetStickerSet
	DeleteSticker
)

type Sticker struct {
	r   Repository
	log *slog.Logger
}

// Sticker represents a sticker.
func New(r Repository, log *slog.Logger) *Sticker {
	return &Sticker{
		r:   r,
		log: log,
	}
}
// Prepare data for getting sticker set.
func (s *Sticker) Prepare(ctx context.Context, msg *tele.Message, b *tele.Bot, code int) error {
	s.log.Info("Prepare.sticker")
	user := user.User{
		ID:       strconv.Itoa(int(msg.Sender.ID)),
		UserName: msg.Sender.Username,
		IsCreate: code == CreateStickerSet,
		IsAdd:    code == AddSticker,
		IsGet:    code == GetStickerSet,
		IsDelete: code == DeleteSticker,
		IterAction: 1,
	}
	// Serialize the user.
	err := user.Serialize()
	if err != nil {
		s.log.Error("Prepare. serialize: %v", err)
		return err
	}
	// Add the user to the repository.
	if err := s.r.Create(ctx, user.ID, user.ByteUser); err != nil {
		s.log.Error("Prepare. add user: %w", err)
		return err
	}
	return nil
}
// Update entity.
func (s *Sticker) Update(ctx context.Context, msg *tele.Message, b *tele.Bot) (code actions.Actions, err error) {
	s.log.Info("Update.sticker")
	user := user.User{
		ID: strconv.Itoa(int(msg.Sender.ID)),
	}
	// Get the user.
	user.ByteUser, err = s.r.Get(ctx, user.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			s.log.Error(fmt.Sprintf("Update. User %v not found", user.ID))
			return actions.Actions{}, ErrEntityNotFound
		}
		s.log.Error(fmt.Sprintf("Update: %v", err))
		return actions.Actions{}, err
	}
	// Deserialize the user.
	err = user.Deserialize()
	if err != nil {
		s.log.Error(fmt.Sprintf("Update: deserialize: %v", err))
		return actions.Actions{}, err
	}
	s.log.Debug("Sticker Set")
	
	s.log.Debug("Its ok")
	// Detect the action.
	if user.IsCreate || user.IsAdd {
		if user.IterAction == 1 {
			user.SetName = msg.Text
		} else if user.IterAction >= 2{
			user.Emoji = msg.Text
		}
		
	}else if user.IsGet || user.IsDelete {

		user.SetName = msg.Text
	} else{
		s.log.Error(fmt.Sprintf("Update. invalid payload: %s", msg.Text))
        return actions.Actions{}, ErrInvalidPayload
	}
	code.Iter = user.IterAction
	user.IterAction++
	// Serialize the user.
	err = user.Serialize()
	if err != nil {
		s.log.Error(fmt.Sprintf("Update. serialize: %v", err))
		return actions.Actions{}, err
	}
	// Update the user in the repository.
	if err := s.r.Update(ctx, user.ID, user.ByteUser); err != nil {
		s.log.Error(fmt.Sprintf("update user: %v", err))
		return actions.Actions{}, err
	}
	// Return the code of action.
	if user.IsCreate{
		code.Code = CreateStickerSet
	}else if user.IsAdd {
		code.Code = AddSticker
	}else if user.IsGet {
		code.Code = GetStickerSet
	}else if user.IsDelete {
		// Get sticker set. If the set not exist - return error.
		st, err := b.StickerSet(prepareName(msg.Text, user.UserName, b.Me.Username))
		if err != nil {
			s.log.Error(fmt.Sprintf("Get. get sticker set: %v", err))
			return actions.Actions{}, err
		}
		// Send a sticker from the set.
		b.Send(&user, &st.Stickers[len(st.Stickers)-1])
		code.Code = DeleteSticker
	}else{
		s.log.Error("update. empty action")
		return actions.Actions{}, ErrInternalError
	}
	
	return code, nil
}
// Done executing the action with the given parameters.
func (s *Sticker) Done(ctx context.Context, msg *tele.Message, b *tele.Bot) (code int, err error) {
	s.log.Info("Done.sticker")
	// Create the user.
	user := user.User{
		ID: strconv.Itoa(int(msg.Sender.ID)),
	}
	// Get the user.
	user.ByteUser, err = s.r.Get(ctx, user.ID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			s.log.Error(fmt.Sprintf("Done. User %v not found.", user.ID))
			return 0, ErrEntityNotFound
		}
		s.log.Error(fmt.Sprintf("Done: %v", err))
		return 0, err
	}
	// Deserialize the user.
	err = user.Deserialize()
	if err != nil {
		s.log.Error("Done. deserialize: %v", err)
		return 0, err
	}
	// Cheking of all data.
	if err = CheckUserIDandName(&user); err != nil {
		s.log.Error(fmt.Sprintf("Done: %v", err))
        return 0, err
    }
	// choose the actions.
	if user.IsCreate{
		code = CreateStickerSet

		if err := s.create(ctx, msg, b, &user); err != nil {
			s.log.Error(fmt.Sprintf("Done. %v", err))
            return 0, err
        }
	}else if user.IsAdd{
		code = AddSticker

		if err := s.add(msg, b, &user); err != nil {
			s.log.Error(fmt.Sprintf("Done. %v", err))
            return 0, err
        }
	}else if user.IsGet{
		code = GetStickerSet
		
		if err := s.get(ctx, b, &user); err != nil {
			s.log.Error(fmt.Sprintf("Done. %v", err))
            return 0, err
        }
	}else if user.IsDelete{
		code = DeleteSticker

		if err := s.delete(ctx, msg, b, &user); err != nil {
			s.log.Error(fmt.Sprintf("Done. %v", err))
            return 0, err
        }
	}
	return code, nil
}
// Create sticker set, add sticker to set.
func (s *Sticker) create(ctx context.Context, msg *tele.Message, b *tele.Bot, user *user.User) (err error) {
	s.log.Info("Create.sticker")
	//Checking fields from user.
	if err = CheckEmojiandSetName(user); err != nil {
		s.log.Error(fmt.Sprintf("Create: %v", err))
        return err
    }
	// Prepare the parameters for resizing a photo.
	photo := file.Photo{
		Width: 512,
		Height: 512,
	}
	// Download the file on disk.
	photoLocal, err := photo.DownloadPhoto(msg, b) 
	if err != nil {
		s.log.Error("Create. download photo: %w", err)
		return err
	}
	// Create a set of stickers.
	if err := b.CreateStickerSet(user, tele.StickerSet{
		Type:          tele.StickerRegular,
		Name:          prepareName(user.SetName, user.UserName, b.Me.Username),
		Title:         user.SetName,
		Emojis:        user.Emoji,
		ContainsMasks: false,
		PNG:           &photoLocal.File,
	}); err != nil {
		s.log.Error("create sticker: %w", err)
		return err
	}
	// Get sticker set.
	st, err := b.StickerSet(prepareName(user.SetName, user.UserName, b.Me.Username))
	if err != nil {
		s.log.Error("Create. get sticker set: %w", err)
		return err
	}
	// Delete entity from db.
	if err := s.r.Delete(ctx, user.ID); err != nil {
		s.log.Error("Sticker. delete user: %w", err)
		return err
	}
	// Send a sticker from the sticker set.
	b.Send(user, &st.Stickers[0])
	return nil
}

func (s *Sticker) add(msg *tele.Message, b *tele.Bot, user *user.User) (err error){
	s.log.Info("add.sticker")
	//Checking fields from user.
	if user.SetName == "" || user.Emoji == "" {
		s.log.Error(fmt.Sprintf("Add. user %s not found.", user.ID))
		return ErrEntityNotFound
	}
	// Prepare the parameters for resizing a photo.
	photo := file.Photo{
		Width: 512,
		Height: 512,
	}
	// Download the file on disk.
	photoLocal, err := photo.DownloadPhoto(msg, b) 
	if err != nil {
		s.log.Error("Create. download photo: %w", err)
		return err
	}
	// Add a sticker to set by photo.
	if err := b.AddSticker(user, tele.StickerSet{
		Name:   prepareName(user.SetName, user.UserName, b.Me.Username),
		Emojis: user.Emoji,
		PNG:    &photoLocal.File,
	}); err != nil {
		s.log.Error("add. sticker: %w", err)
		return err
	}
	// Get sticker set.
	st, err := b.StickerSet(prepareName(user.SetName, user.UserName, b.Me.Username))
	if err != nil {
		s.log.Error("Get. get sticker set: %w", err)
		return err
	}
	// Send a sticker from the set.
	b.Send(user, &st.Stickers[len(st.Stickers)-1])
	return nil
}
// Get sticker set.
func (s *Sticker) get(ctx context.Context, b *tele.Bot, user *user.User) (err error) {
	s.log.Info("Get.sticker")
	//Checking fields from user.
	if err = CheckSetName(user); err != nil {
		s.log.Error("Get. %w", err)
		return err
	}
	// Get sticker set.
	st, err := b.StickerSet(prepareName(user.SetName, user.UserName, b.Me.Username))
	if err != nil {
		s.log.Error("Get. get sticker set: %w", err)
		return err
	}
	if len(st.Stickers) == 0 {
		s.log.Error("Get. sticker set not found.")
        return ErrStickersIsDeleted
	}
	// Delete entity from db.
	if err := s.r.Delete(ctx, user.ID); err != nil {
		s.log.Error("Sticker. delete user: %w", err)
		return err
	}
	// Send a sticker from the set.
	b.Send(user, &st.Stickers[0])
	return nil
}
// Delete a sticker.
func (s *Sticker) delete(ctx context.Context, msg *tele.Message, b *tele.Bot, user *user.User) (err error) {
	s.log.Info("Sticker. delete sticker")
	//Checking fields from user.
	if err = CheckSetName(user); err != nil {
		s.log.Error("Delete. %w", err)
		return err
	}
	// Delete a sticker.
	if err := b.DeleteSticker(msg.Sticker.FileID); err != nil {
		s.log.Error("Delete. delete sticker: %w", err)
		return err
	}
	// Delete entity from db.
	if err := s.r.Delete(ctx, user.ID); err != nil {
		s.log.Error("Delete. delete user: %w", err)
		return err
	}
	return nil
}

func prepareName(stickerSetName, userName, botName string) string {
	return stickerSetName + "_" + userName + "_by_" + botName
}
