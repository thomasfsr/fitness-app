package http

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/yourusername/fitness-app/internal/domain/message"
    "github.com/yourusername/fitness-app/internal/usecase"
)

type messageRequest struct {
    UserId  uint32 `json:"user_id" binding:"required"`
    Role    string `json:"role" binding:"required"`
    Message string `json:"message" binding:"required"`
}

func NewMessageHandler(rg *gin.RouterGroup, uc *usecase.MessageUseCase) {
    h := &messageHandler{uc: uc}
    m := rg.Group("/messages")
    m.POST("", h.Create)
    m.GET("/user/:user_id", h.ListByUser)
}

type messageHandler struct {
    uc *usecase.MessageUseCase
}

func (h *messageHandler) Create(c *gin.Context) {
    var req messageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    mm := &message.Message{
        UserId:  req.UserId,
        Role:    req.Role,
        Message: req.Message,
    }
    if err := h.uc.CreateMessage(mm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *messageHandler) ListByUser(c *gin.Context) {
    idS := c.Param("user_id")
    id64, _ := strconv.ParseUint(idS, 10, 32)
    msgs, err := h.uc.ListByUser(uint32(id64))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, msgs)
}
