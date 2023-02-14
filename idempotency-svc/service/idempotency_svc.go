package service

import (
	"context"
	"fmt"
	"max-idempotency-svc/model"

	"github.com/go-redis/redis/v8"
)

type idempotencySvc struct {
	redisClient *redis.Client
}

func NewIdempotencySvc() *idempotencySvc {
	opts := &redis.Options{
		Addr:     "redis_db:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(opts)
	return &idempotencySvc{
		redisClient: client,
	}
}

// Check interface complience
var _ IdempotencyService = (*idempotencySvc)(nil)

func (s *idempotencySvc) Validate(ctx context.Context, t model.Transaction) (string, error) {
	cmd := s.redisClient.Get(ctx, t.ID)
	cmderr := cmd.Err()
	if cmderr == redis.Nil {
		err := s.redisClient.Set(ctx, t.ID, t.Timestamp, 0).Err()
		if err != nil {
			return "", fmt.Errorf("[Rdis Validate] %s", err)
		}
		return "", nil
	}

	if cmderr != nil {
		return "", fmt.Errorf("[Rdis Validate] %s", cmderr)
	}

	val, err := cmd.Result()
	if err != nil {
		return "", fmt.Errorf("[Rdis Validate] Transaction duplicate")
	}

	return fmt.Sprint("Result:", val), nil
}
