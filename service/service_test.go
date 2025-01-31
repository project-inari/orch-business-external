package service

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/project-inari/orch-business-external/config"
	"github.com/project-inari/orch-business-external/dto"
	"github.com/project-inari/orch-business-external/pkg/httpclient"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

type mockAuthMiddlewareRepository struct {
	err error
}

func (m *mockAuthMiddlewareRepository) VerifyToken(_ context.Context, _ string) error {
	return m.err
}

type mockAPICoreBusinessServerRepository struct {
	createRes *httpclient.Response[dto.CoreBusinessServerCreateRes]
	err       error
}

func (m *mockAPICoreBusinessServerRepository) CallCreate(_ context.Context, _ dto.CoreBusinessServerCreateReq) (*httpclient.Response[dto.CoreBusinessServerCreateRes], error) {
	return m.createRes, m.err
}

type mockCacheRepository struct {
	getRes *redis.StringCmd
	setRes *redis.StatusCmd
}

func (m *mockCacheRepository) Get(_ context.Context, _ string) *redis.StringCmd {
	return m.getRes
}

func (m *mockCacheRepository) Set(_ context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return m.setRes
}

const (
	mockName             = "mockName"
	mockIndustryType     = "mockIndustryType"
	mockBusinessType     = "mockBusinessType"
	mockDescription      = "mockDescription"
	mockPhoneNo          = "mockPhoneNo"
	mockOperatingHours   = `{"monday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"tuesday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"wednesday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"thursday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"friday":{"open":true,"openTime":"09:00","closeTime":"17:00"},"saturday":{"open":false},"sunday":{"open":false}}`
	mockAddress          = "mockAddress"
	mockBusinessImageURL = "mockBusinessImageURL"
	mockOwnerUsername    = "mockOwnerUsername"
)

var (
	mockOperatingHoursStruct = dto.OperatingHours{
		Monday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Tuesday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Wednesday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Thursday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Friday: dto.OpenTime{
			Open:      true,
			OpenTime:  "09:00",
			CloseTime: "17:00",
		},
		Saturday: dto.OpenTime{
			Open: false,
		},
		Sunday: dto.OpenTime{
			Open: false,
		},
	}
)

func TestCreateNewBusiness(t *testing.T) {
	ctx := context.Background()

	mockAuthMiddlewareRepository := &mockAuthMiddlewareRepository{
		err: nil,
	}

	req := dto.CreateNewBusinessReq{
		Name:             mockName,
		IndustryType:     mockIndustryType,
		BusinessType:     mockBusinessType,
		Description:      mockDescription,
		PhoneNo:          mockPhoneNo,
		OperatingHours:   mockOperatingHoursStruct,
		Address:          mockAddress,
		BusinessImageURL: mockBusinessImageURL,
		OwnerUsername:    mockOwnerUsername,
	}

	expectedRes := &dto.CreateNewBusinessRes{
		BusinessID:   1,
		BusinessName: mockName,
		Success:      true,
	}

	t.Run("success", func(t *testing.T) {
		mockAPICoreBusinessServerRepository := &mockAPICoreBusinessServerRepository{
			createRes: &httpclient.Response[dto.CoreBusinessServerCreateRes]{
				HTTPStatusCode: http.StatusOK,
				Response: dto.CoreBusinessServerCreateRes{
					BusinessID:   1,
					BusinessName: mockName,
					Success:      true,
				},
			},
		}

		mockCacheRepository := &mockCacheRepository{}

		svc := New(Dependencies{
			Conf:                            &config.Config{},
			AuthMiddlewareRepository:        mockAuthMiddlewareRepository,
			CacheRepository:                 mockCacheRepository,
			APICoreBusinessServerRepository: mockAPICoreBusinessServerRepository,
		})

		res, err := svc.CreateNewBusiness(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("error - when call api create core-business-server httpclient returned error", func(t *testing.T) {
		mockAPICoreBusinessServerRepository := &mockAPICoreBusinessServerRepository{
			err: assert.AnError,
		}

		mockCacheRepository := &mockCacheRepository{}

		svc := New(Dependencies{
			Conf:                            &config.Config{},
			AuthMiddlewareRepository:        mockAuthMiddlewareRepository,
			CacheRepository:                 mockCacheRepository,
			APICoreBusinessServerRepository: mockAPICoreBusinessServerRepository,
		})

		res, err := svc.CreateNewBusiness(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("error - when call api create core-business-server httpclient returned http status code not 200", func(t *testing.T) {
		mockAPICoreBusinessServerRepository := &mockAPICoreBusinessServerRepository{
			createRes: &httpclient.Response[dto.CoreBusinessServerCreateRes]{
				HTTPStatusCode: http.StatusBadRequest,
				Response: dto.CoreBusinessServerCreateRes{
					BusinessID:   0,
					BusinessName: "",
					Success:      false,
				},
			},
		}

		mockCacheRepository := &mockCacheRepository{}

		svc := New(Dependencies{
			Conf:                            &config.Config{},
			AuthMiddlewareRepository:        mockAuthMiddlewareRepository,
			CacheRepository:                 mockCacheRepository,
			APICoreBusinessServerRepository: mockAPICoreBusinessServerRepository,
		})

		res, err := svc.CreateNewBusiness(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
