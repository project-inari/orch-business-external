package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"

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

	key := fmt.Sprintf("%s:%s", s.conf.RedisConfig.KeyUserVerifiedAccount, req.OwnerUsername)
	currentSession, err := s.getUserSession(ctx, key)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err != redis.Nil {
		currentSession.Businesses = append(currentSession.Businesses, dto.Businesses(req))

		cacheRes := s.cacheRepository.SetWithRemainingTTL(ctx, key, currentSession)
		if cacheRes.Err() != nil {
			return nil, cacheRes.Err()
		}
	}

	res := dto.CreateNewBusinessRes(coreRes.Response)

	return &res, nil
}

func (s *service) getUserSession(ctx context.Context, key string) (*dto.UserVerifiedAccountCache, error) {
	var session dto.UserVerifiedAccountCache
	result, err := s.cacheRepository.Get(ctx, key).Result()
	if err != nil {
		slog.Error("get cache: " + err.Error())
		return nil, err
	}

	if err := json.NewDecoder(bytes.NewBufferString(result)).Decode(&session); err != nil {
		slog.Error("decode cache: " + err.Error())
		return nil, err
	}
	return &session, nil
}
