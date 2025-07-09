package http

import (
	"net/http"

	"ragnarok-multifinance/internal/model/request"
	"ragnarok-multifinance/internal/model/response"
	"ragnarok-multifinance/internal/pkg"
	"ragnarok-multifinance/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type TransactionHandler struct {
	uc          *usecase.TransactionUsecase
	redisClient *redis.Client
}

func NewTransactionHandler(uc *usecase.TransactionUsecase, redisClient *redis.Client) *TransactionHandler {
	return &TransactionHandler{uc: uc, redisClient: redisClient}
}

func (h *TransactionHandler) Register(r *gin.RouterGroup) {
	r.POST("/transactions", h.Create)
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req request.TransactionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := pkg.GetUserID(c)
	if userID == "" || userID != req.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user session"})
		return
	}
	ctx := c.Request.Context()
	sessionKey := "SESSION:" + userID
	if err := h.redisClient.Get(ctx, sessionKey).Err(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired or not found"})
		return
	}
	tx, err := h.uc.Create(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := response.TransactionResponse{
		ID:                tx.ID,
		UserID:            tx.UserID,
		ProductID:         tx.ProductID,
		OTR:               tx.OTR,
		AdminFee:          tx.AdminFee,
		InstallmentAmount: tx.InstallmentAmount,
		AssetName:         tx.AssetName,
		Status:            tx.Status,
		CreatedAt:         tx.CreatedAt,
	}
	c.JSON(http.StatusCreated, resp)
}
