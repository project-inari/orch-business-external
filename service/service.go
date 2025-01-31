// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/project-inari/orch-business-external/config"
	"github.com/project-inari/orch-business-external/dto"
	"github.com/project-inari/orch-business-external/repository"
)

// Port represents the service layer functions
type Port interface {
	Auth(c echo.Context, token string) (echo.Context, error)
	CreateNewBusiness(ctx context.Context, req dto.CreateNewBusinessReq) (*dto.CreateNewBusinessRes, error)
}

type service struct {
	conf                            *config.Config
	authMiddlewareRepository        repository.AuthMiddlewareRepository
	cacheRepository                 repository.CacheRepository
	apiCoreBusinessServerRepository repository.APICoreBusinessServerRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	Conf                            *config.Config
	AuthMiddlewareRepository        repository.AuthMiddlewareRepository
	CacheRepository                 repository.CacheRepository
	APICoreBusinessServerRepository repository.APICoreBusinessServerRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		conf:                            d.Conf,
		authMiddlewareRepository:        d.AuthMiddlewareRepository,
		cacheRepository:                 d.CacheRepository,
		apiCoreBusinessServerRepository: d.APICoreBusinessServerRepository,
	}
}
