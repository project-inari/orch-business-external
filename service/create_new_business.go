package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/project-inari/orch-business-external/dto"
)

func (s *service) CreateNewBusiness(ctx context.Context, req dto.CreateNewBusinessReq) (*dto.CreateNewBusinessRes, error) {
	coreRes, err := s.apiCoreBusinessServerRepository.CallCreate(ctx, dto.CoreBusinessServerCreateReq(req))
	if err != nil {
		return nil, err
	}
	if coreRes.HTTPStatusCode != http.StatusOK {
		return nil, fmt.Errorf("error calling core-business-server returned status code not 200: %v", coreRes.HTTPStatusCode)
	}

	res := dto.CreateNewBusinessRes(coreRes.Response)

	return &res, nil
}
