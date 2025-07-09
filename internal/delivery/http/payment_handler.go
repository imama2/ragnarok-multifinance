package http

import (
	"net/http"

	"ragnarok-multifinance/internal/model/database"
	"ragnarok-multifinance/internal/model/request"
	"ragnarok-multifinance/internal/model/response"
	"ragnarok-multifinance/internal/pkg"
	"ragnarok-multifinance/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type PaymentHandler struct {
	uc          *usecase.PaymentUsecase
	redisClient *redis.Client
}

func NewPaymentHandler(uc *usecase.PaymentUsecase, redisClient *redis.Client) *PaymentHandler {
	return &PaymentHandler{uc: uc, redisClient: redisClient}
}

func (h *PaymentHandler) Register(r *gin.RouterGroup) {
	r.POST("/payments", h.RecordPayment)
	r.GET("/transactions/:id/schedule", h.GetSchedule)
}

func (h *PaymentHandler) RecordPayment(c *gin.Context) {
	var req request.PaymentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := pkg.GetUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user session"})
		return
	}
	ctx := c.Request.Context()
	sessionKey := "SESSION:" + userID
	if err := h.redisClient.Get(ctx, sessionKey).Err(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired or not found"})
		return
	}
	payment := &database.Payment{
		TransactionID: req.TransactionID,
		Amount:        req.Amount,
	}
	if err := h.uc.RecordPayment(ctx, payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "payment recorded"})
}

func (h *PaymentHandler) GetSchedule(c *gin.Context) {
	transactionID := c.Param("id")
	ctx := c.Request.Context()
	payments, err := h.uc.GetSchedule(ctx, transactionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var resp []response.PaymentResponse
	for _, p := range payments {
		resp = append(resp, response.PaymentResponse{
			ID:            p.ID,
			TransactionID: p.TransactionID,
			DueDate:       p.DueDate,
			Amount:        p.Amount,
			Status:        p.Status,
			PaidAt:        p.PaidAt,
		})
	}
	c.JSON(http.StatusOK, resp)
}
