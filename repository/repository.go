// Package repository provides the repository interfaces for the domain
package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/project-inari/orch-business-external/dto"
	"github.com/project-inari/orch-business-external/pkg/httpclient"
)

// AuthMiddlewareRepository represents the repository layer functions of auth middleware repository
type AuthMiddlewareRepository interface {
	VerifyToken(ctx context.Context, token string) error
}

// CacheRepository represents the repository layer functions of cache repository
type CacheRepository interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) *redis.StatusCmd
	SetWithRemainingTTL(ctx context.Context, key string, value interface{}) *redis.StatusCmd
}

// APICoreBusinessServerRepository represents the repository layer functions of api core business server repository
type APICoreBusinessServerRepository interface {
	CallCreate(ctx context.Context, req dto.CoreBusinessServerCreateReq) (*httpclient.Response[dto.CoreBusinessServerCreateRes], error)
}
