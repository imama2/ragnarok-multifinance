package http

import (
	"net/http"

	"ragnarok-multifinance/internal/model/request"
	"ragnarok-multifinance/internal/model/response"
	"ragnarok-multifinance/internal/pkg"
	"ragnarok-multifinance/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) Register(r *gin.RouterGroup) {
	r.GET("/products", h.List)
	r.POST("/products", h.Create)
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.uc.List(c.Request.Context())
	if err != nil {
		logrus.Error("List products error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list products"})
		return
	}
	var resp []response.ProductResponse
	for _, p := range products {
		resp = append(resp, response.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Tenure:      p.Tenure,
			MaxLimit:    p.MaxLimit,
			Interest:    p.Interest,
			Description: p.Description,
			IsActive:    p.IsActive,
		})
	}
	c.JSON(http.StatusOK, response.ProductListResponse{Products: resp})
}

func (h *ProductHandler) Create(c *gin.Context) {
	// Enforce admin-only access
	if pkg.GetUserRole(c) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		return
	}
	var req request.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := h.uc.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := response.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Tenure:      product.Tenure,
		MaxLimit:    product.MaxLimit,
		Interest:    product.Interest,
		Description: product.Description,
		IsActive:    product.IsActive,
	}
	c.JSON(http.StatusCreated, resp)
}
