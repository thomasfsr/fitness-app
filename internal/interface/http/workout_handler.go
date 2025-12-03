package http

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/thomasfsr/fitness-app/internal/domain/workout"
    "github.com/thomasfsr/fitness-app/internal/usecase"
)

type workoutRequest struct {
    UserId   uint32 `json:"user_id" binding:"required"`
    Exercise string `json:"exercise" binding:"required"`
    Weight   uint16 `json:"weight"`
    Reps     uint8  `json:"reps"`
}

func NewWorkoutHandler(rg *gin.RouterGroup, uc *usecase.WorkoutUseCase) {
    h := &workoutHandler{uc: uc}
    w := rg.Group("/workouts")
    w.POST("", h.Create)
    w.GET("/user/:user_id", h.ListByUser)
    w.GET("/:id", h.Get)
    w.PUT("/:id", h.Update)
    w.DELETE("/:id", h.Delete)
}

type workoutHandler struct {
    uc *usecase.WorkoutUseCase
}

func (h *workoutHandler) Create(c *gin.Context) {
    var req workoutRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    w := &workout.WorkoutSet{
        UserId:   req.UserId,
        Exercise: req.Exercise,
        Weight:   req.Weight,
        Reps:     req.Reps,
    }
    if err := h.uc.CreateWorkout(w); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *workoutHandler) Get(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 64)
    w_, err := h.uc.GetWorkout(id64)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, w_)
}

func (h *workoutHandler) Update(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 64)
    var req workoutRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    w := &workout.WorkoutSet{
        ID:       id64,
        UserId:   req.UserId,
        Exercise: req.Exercise,
        Weight:   req.Weight,
        Reps:     req.Reps,
    }
    if err := h.uc.UpdateWorkout(w); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *workoutHandler) Delete(c *gin.Context) {
    idS := c.Param("id")
    id64, _ := strconv.ParseUint(idS, 10, 64)
    if err := h.uc.DeleteWorkout(id64); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *workoutHandler) ListByUser(c *gin.Context) {
    idS := c.Param("user_id")
    id64, _ := strconv.ParseUint(idS, 10, 32)
    ws, err := h.uc.ListByUser(uint32(id64))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ws)
}
