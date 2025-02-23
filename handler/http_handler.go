package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/project-inari/orch-business-external/dto"
	"github.com/project-inari/orch-business-external/pkg/request"
	"github.com/project-inari/orch-business-external/pkg/response"
)

const (
	errorCodeRequestPayloadInvalid = "IERC10000"
	errorCodeMissingToken          = "IERC10005"
	errorCodeInvalidToken          = "IERC10006"
	errorCodeDefaultError          = "IERC50000"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

func (h *httpHandler) CreateNewBusiness(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewBusinessReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewBusiness] bad request: %v", err), errorCodeRequestPayloadInvalid)
	}

	res, err := h.d.Service.CreateNewBusiness(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewBusiness] internal server error: %v", err), errorCodeDefaultError)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
