package RedisImplementation

import (
	"SimpleRedis/Model"
	"SimpleRedis/RedisService"
	"context"
	"encoding/json"
	"strconv"
	"github.com/redis/go-redis/v9"
)

var (
	keyUser = "user"
)


func NewRedisService(address string, password string, db int) (RedisService.RedisService, *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &redisServiceImplement{client}, client
}

type redisServiceImplement struct {
	RedisClient *redis.Client
}

func (r *redisServiceImplement) get(ctx context.Context, key string, expectedOutput interface{}) (error) {


	val, err := r.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(val, expectedOutput)
}

func (r *redisServiceImplement) set(ctx context.Context, key string, value interface{}) (error) {
	

	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.RedisClient.Set(ctx, key, val, 0).Err()
}

func (r *redisServiceImplement) delete(ctx context.Context, key string) (error) {
	

	return r.RedisClient.Del(ctx, key).Err()
}

func (r *redisServiceImplement) GetUser(ctx context.Context, id uint32) (*Model.User, error) {
	var user Model.User
	key := keyUser + ":" + strconv.Itoa(int(id))

	err := r.get(ctx, key, &user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (r *redisServiceImplement) SetUser(ctx context.Context,user *Model.User) error {
	key := keyUser + ":" + strconv.Itoa(int(user.ID))
	return r.set(ctx, key, user)
}
	
func (r *redisServiceImplement) DeleteUser(ctx context.Context, id uint32) error {
	key := keyUser + ":" + strconv.Itoa(int(id))
	return r.delete(ctx, key)
}