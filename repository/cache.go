package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	client                 *redis.Client
	keyUserVerifiedAccount string
}

// CacheRepositoryConfig represents the configuration of the cache repository
type CacheRepositoryConfig struct {
	KeyUserVerifiedAccount string
}

// CacheRepositoryDependencies represents the dependencies of the cache repository
type CacheRepositoryDependencies struct {
	Client *redis.Client
}

// NewCacheRepository creates a new cache repository
func NewCacheRepository(c CacheRepositoryConfig, d CacheRepositoryDependencies) CacheRepository {
	return &cacheRepository{
		client:                 d.Client,
		keyUserVerifiedAccount: c.KeyUserVerifiedAccount,
	}
}

func (r *cacheRepository) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *cacheRepository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd {
	marshalValue, _ := json.Marshal(value)
	return r.client.Set(ctx, key, marshalValue, ttl)
}
