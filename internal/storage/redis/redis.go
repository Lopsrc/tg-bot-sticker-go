package redis

import (
	"context"
	"fmt"

	"tg-bot-sticker-go/internal/config"
	"tg-bot-sticker-go/internal/storage"

	"github.com/redis/go-redis/v9"
)

type Storage struct{
	Client *redis.Client
}
// New creates a new Redis storage instance.
func New(cfg *config.Config) *Storage {
	client := redis.NewClient(&redis.Options{
        Addr:	  preparePath(cfg.Storage.Hostname, cfg.Storage.Port),
        Password: cfg.Storage.Password,
        DB:		  0,  // use default DB
    })

	pong, err := client.Ping(context.Background()).Result()
	if err != nil{
		fmt.Printf("err redis: %v\n", err)
		panic(pong)
	}
	return &Storage{
		Client: client,
	}
}
// Create adds a new user to the storage.
func (s *Storage) Create(ctx context.Context, chatID string, user []byte) error{

	err := s.Client.Set(ctx, chatID, user, 0).Err()
	if err!= nil {
        return err
    }
	return nil
}
// Get retrieves the user data
func (s *Storage) Get(ctx context.Context, chatID string) (user []byte, err error) {
	user, err = s.Client.Get(ctx, chatID).Bytes()
	if err != nil {
		return nil, storage.ErrNotFound
	}
	return user, err
}
// Update updates the user data
func (s *Storage) Update(ctx context.Context, chatID string, user []byte) error{

	err := s.Client.Set(ctx, chatID, user, 0).Err()
	if err!= nil {
        return err
    }
    return nil
}
// IsExist returns 1 if user exists and 0 otherwise.
func (s *Storage) IsExist(ctx context.Context, chatID string) (int64, error) { 

	res, err := s.Client.Exists(ctx, chatID).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}
// Delete removes the user data from the storage.
func (s *Storage) Delete(ctx context.Context, chatID string) error{
	
	err := s.Client.Del(ctx, chatID).Err()
    if err!= nil {
        return err
    }
    return nil
}

func preparePath(localhost string, port string) string{
	return localhost + ":" + port
}