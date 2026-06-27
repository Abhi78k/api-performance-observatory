package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type EndpointHandler struct {
	endpointService *services.EndpointService
}

func NewEndpointHandler(
	endpointService *services.EndpointService,
) *EndpointHandler {

	return &EndpointHandler{
		endpointService: endpointService,
	}
}

// CreateEndpoint godoc
//
// @Summary Create endpoint
// @Description Creates a new endpoint to monitor.
// @Tags Endpoints
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param request body dto.CreateEndpointRequest true "Endpoint details"
//
// @Success 201 {object} dto.EndpointSuccessResponse
// @Failure 400 {object} utils.ValidationErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /endpoints [post]
func (h *EndpointHandler) CreateEndpoint(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request.")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationError(
			c,
			utils.FormatValidationErrors(err),
		)
		return
	}

	userID := c.MustGet("UserID").(uint)

	endpoint, err := h.endpointService.CreateEndpoint(
		ctx,
		req.Name,
		req.URL,
		req.ExpectedStatus,
		userID,
	)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.Created(c, endpoint)
}

// GetEndpoints godoc
//
// @Summary List endpoints
// @Description Returns all monitored endpoints.
// @Tags Endpoints
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.EndpointListResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /endpoints [get]
func (h *EndpointHandler) GetEndpoints(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.MustGet("UserID").(uint)
	page, limit := utils.GetPaginationParams(c)
	search := c.Query("search")
	status := c.Query("status")

	endpoints, total, err := h.endpointService.GetEndpointsPaginated(ctx, userID, page, limit, search, status)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.PaginatedOK(c, endpoints, page, limit, total)
}

// GetEndpoint godoc
//
// @Summary Get endpoint
// @Description Returns one endpoint.
// @Tags Endpoints
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.EndpointSuccessResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
//
// @Router /endpoints/{id} [get]
func (h *EndpointHandler) GetEndpoint(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	endpoint, err := h.endpointService.GetEndpoint(
		ctx,
		id,
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.OK(c, endpoint)
}

// UpdateEndpoint godoc
//
// @Summary Update endpoint
// @Description Updates an existing endpoint.
// @Tags Endpoints
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
// @Param request body dto.UpdateEndpointRequest true "Updated endpoint"
//
// @Success 200 {object} dto.EndpointSuccessResponse
// @Failure 400 {object} utils.ValidationErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /endpoints/{id} [put]
func (h *EndpointHandler) UpdateEndpoint(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	var req dto.UpdateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body.")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationError(
			c,
			utils.FormatValidationErrors(err),
		)
		return
	}

	endpoint, err := h.endpointService.UpdateEndpoint(
		ctx,
		id,
		userID,
		req.Name,
		req.URL,
		req.ExpectedStatus,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.OK(c, endpoint)
}

// DeleteEndpoint godoc
//
// @Summary Delete endpoint
// @Description Deletes an endpoint.
// @Tags Endpoints
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.MessageResponse
// @Failure 404 {object} utils.ErrorResponse
//
// @Router /endpoints/{id} [delete]
func (h *EndpointHandler) DeleteEndpoint(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	err := h.endpointService.DeleteEndpoint(
		ctx,
		id,
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.Message(c, http.StatusOK, "Endpoint deleted successfully.")
}
