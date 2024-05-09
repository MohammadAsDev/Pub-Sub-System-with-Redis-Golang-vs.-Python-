package publisher

import (
	"context"
	"errors"

	"github.com/MohammadAsDev/pub_sub/config"
	"github.com/redis/go-redis/v9"
)

type RedisPublisher struct {
	_Id      int
	_Client  *redis.Client
	_Context context.Context
	_Channel string
}

var publishersCounter = 0

func NewRedisPublisher(ctx context.Context, channel string) (Publisher, error) {

	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	redis_config := config.RedisConfig

	publishersCounter += 1
	client := redis.NewClient(&redis.Options{
		Addr:     redis_config.Addr,
		DB:       0,
		Password: redis_config.Password,
	})

	if client == nil {
		return nil, errors.New("[publisher error]: can't create redis client, check connection configuration")
	}

	return &RedisPublisher{
		_Id:      publishersCounter,
		_Client:  client,
		_Context: ctx,
		_Channel: channel,
	}, nil
}

func (publisher *RedisPublisher) Publish(message string) error {
	_, err := publisher._Client.Publish(publisher._Context, publisher._Channel, message).Result()
	return err
}

func (publisher *RedisPublisher) Id() int {
	return publisher._Id
}
