package RedisService

import (
	"SimpleRedis/Model"
	"context"
)

type RedisService interface {
	GetUser(ctx context.Context, id uint32) (*Model.User, error)
	SetUser(ctx context.Context, user *Model.User) error
	DeleteUser(ctx context.Context, id uint32) error
}

