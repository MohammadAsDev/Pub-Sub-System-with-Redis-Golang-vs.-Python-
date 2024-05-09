package subscriber

import (
	"context"
	"errors"

	"github.com/MohammadAsDev/pub_sub/config"
	"github.com/redis/go-redis/v9"
)

var subscribersCounter = 0

const timeOutSecs = 5

type RedisSubscriber struct {
	_Id      int
	_Client  *redis.Client
	_Context context.Context
	_Channel string
}

func NewRedisSubscriber(ctx context.Context, channel string) (Subscriber, error) {

	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	redis_config := config.RedisConfig

	client := redis.NewClient(&redis.Options{
		Addr:     redis_config.Addr,
		DB:       0,
		Password: redis_config.Password,
	})

	if client == nil {
		return nil, errors.New("[subscriber error]: can't create redis client, check connection configuration")
	}

	subscribersCounter += 1
	return &RedisSubscriber{
		_Id:      subscribersCounter,
		_Client:  client,
		_Context: ctx,
		_Channel: channel,
	}, nil
}

func (subscriber *RedisSubscriber) Subscribe() (chan string, chan error) {

	messages_chan := make(chan string)
	errs_chan := make(chan error)
	pubsub := subscriber._Client.Subscribe(subscriber._Context, subscriber._Channel)

	go func() {
		for {
			message, err := pubsub.ReceiveMessage(subscriber._Context)
			messages_chan <- message.Payload
			if err != nil {
				errs_chan <- err
			}
		}
	}()

	return messages_chan, errs_chan
}

func (subscriber *RedisSubscriber) Id() int {
	return subscriber._Id
}
