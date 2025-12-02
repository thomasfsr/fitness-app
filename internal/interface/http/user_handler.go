package http

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/yourusername/fitness-app/internal/domain/user"
    "github.com/yourusername/fitness-app/internal/usecase"
)

type userRequest struct {
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
    Whatsapp  uint64 `json:"whatsapp"`
    Active    bool   `json:"active"`
}

func NewUserHandler(rg *gin.RouterGroup, uc *usecase.UserUseCase) {
    h := &userHandler{uc: uc}
    users := rg.Group("/users")
    users.POST("", h.Create)
    users.GET("", h.List)
    users.GET("/:id", h.Get)
    users.PUT("/:id", h.Update)
    users.DELETE("/:id", h.Delete)
}

type userHandler struct {
    uc *usecase.UserUseCase
}

func (h *userHandler) Create(c *gin.Context) {
    var req userRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    u := &user.User{
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Whatsapp:  req.Whatsapp,
        Active:    req.Active,
    }
    if err := h.uc.CreateUser(u); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *userHandler) Get(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 32)
    u_, err := h.uc.GetUser(uint32(id64))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, u_)
}

func (h *userHandler) Update(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 32)
    var req userRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    u := &user.User{
        ID:        uint32(id64),
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Whatsapp:  req.Whatsapp,
        Active:    req.Active,
    }
    if err := h.uc.UpdateUser(u); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *userHandler) Delete(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 32)
    if err := h.uc.DeleteUser(uint32(id64)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *userHandler) List(c *gin.Context) {
    users, err := h.uc.ListUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}
