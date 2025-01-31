// Package di provides dependency injection for the server
package di

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"github.com/project-inari/orch-business-external/config"
	"github.com/project-inari/orch-business-external/handler"
	"github.com/project-inari/orch-business-external/pkg/httpclient"
	"github.com/project-inari/orch-business-external/repository"
	"github.com/project-inari/orch-business-external/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                c.SentryConfig.SentryDSN,
		Debug:              false,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [main.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// GRPC Auth Client initialization
	grpcClient, err := newGRPCClient(grpcClientOptions{
		url: c.GRPCAuthConfig.URL,
	})
	if err != nil {
		log.Panicf("error - [main.New] unable to create grpc client: %v", err)
	}

	// Redis initialization
	redisClient, err := newRedis(redisOptions{
		host:     c.RedisConfig.Host,
		password: c.RedisConfig.Password,
		timeout:  c.RedisConfig.Timeout,
		maxRetry: c.RedisConfig.MaxRetry,
		poolSize: c.RedisConfig.PoolSize,
		db:       c.RedisConfig.DB,
	})
	if err != nil {
		log.Panicf("error - [main.New] unable to connect to Redis: %v", err)
	}
	defer func() {
		if err := redisClient.client.Close(); err != nil {
			slog.Error("error - [main.New] unable to close Redis connection", slog.Any("error", err))
		}
	}()

	// HTTP Client initialization
	httpClientCoreBusinessServer := httpclient.NewHTTPClient(httpclient.Options{
		MaxConns:                 c.APICoreBusinessServerConfig.MaxConns,
		MaxRetry:                 c.APICoreBusinessServerConfig.MaxRetry,
		Timeout:                  c.APICoreBusinessServerConfig.Timeout,
		InsecureSkipVerify:       c.APICoreBusinessServerConfig.InsecureSkipVerify,
		MaxTransactionsPerSecond: c.APICoreBusinessServerConfig.MaxTransactionsPerSecond,
	})

	// Repository initialization
	authRepo := repository.NewAuthMiddlewareRepository(repository.AuthMiddlewareRepositoryDependencies{
		GRPCClient: grpcClient.client,
	})

	cacheRepo := repository.NewCacheRepository(repository.CacheRepositoryConfig{
		KeyUserVerifiedAccount: c.RedisConfig.KeyUserVerifiedAccount,
	}, repository.CacheRepositoryDependencies{
		Client: redisClient.client,
	})

	apiCoreBusinessServerRepo := repository.NewAPICoreBusinessServerRepository(repository.APICoreBusinessServerRepositoryConfig{
		BaseURL:    c.APICoreBusinessServerConfig.BaseURL,
		CreatePath: c.APICoreBusinessServerConfig.CreatePath,
	}, repository.APICoreBusinessServerRepositoryDependencies{
		Client: httpClientCoreBusinessServer,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		Conf:                            c,
		AuthMiddlewareRepository:        authRepo,
		CacheRepository:                 cacheRepo,
		APICoreBusinessServerRepository: apiCoreBusinessServerRepo,
	})

	// Handler initialization
	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
