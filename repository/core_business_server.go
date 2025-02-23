package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/orch-business-external/dto"
	"github.com/project-inari/orch-business-external/pkg/httpclient"
)

type apiCoreBusinessServerRepository struct {
	baseURL    string
	createPath string
	client     *http.Client
}

// APICoreBusinessServerRepositoryConfig represents the configuration of the api core business server repository
type APICoreBusinessServerRepositoryConfig struct {
	BaseURL    string
	CreatePath string
}

// APICoreBusinessServerRepositoryDependencies represents the dependencies of the api core business server repository
type APICoreBusinessServerRepositoryDependencies struct {
	Client *http.Client
}

// NewAPICoreBusinessServerRepository creates a new api core business server repository
func NewAPICoreBusinessServerRepository(c APICoreBusinessServerRepositoryConfig, d APICoreBusinessServerRepositoryDependencies) APICoreBusinessServerRepository {
	return &apiCoreBusinessServerRepository{
		baseURL:    c.BaseURL,
		createPath: c.CreatePath,
		client:     d.Client,
	}
}

// CallCreate calls the create endpoint of the core business server
func (r *apiCoreBusinessServerRepository) CallCreate(ctx context.Context, req dto.CoreBusinessServerCreateReq) (*httpclient.Response[dto.CoreBusinessServerCreateRes], error) {
	return httpclient.Post[dto.CoreBusinessServerCreateReq, dto.CoreBusinessServerCreateRes](ctx, r.client, fmt.Sprintf("%s%s", r.baseURL, r.createPath), map[string]string{}, req)
}
