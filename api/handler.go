package api

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "go-api-grpc-demo/model"
    "net/http"
)

type Handler struct {
    store *model.UserStore
}

func NewHandler(store *model.UserStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, exists := h.store.Get(id)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *Handler) ListUsers(c *gin.Context) {
    users := h.store.List()
    c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(c *gin.Context) {
    var user model.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user.ID = uuid.New().String()
    h.store.Create(&user)
    c.JSON(http.StatusCreated, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user model.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user.ID = id
    if !h.store.Update(&user) {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if !h.store.Delete(id) {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.Status(http.StatusNoContent)
}

func SetupRouter(handler *Handler) *gin.Engine {
    router := gin.Default()
    api := router.Group("/api")
    {
        users := api.Group("/users")
        {
            users.GET("", handler.ListUsers)
            users.GET("/:id", handler.GetUser)
            users.POST("", handler.CreateUser)
            users.PUT("/:id", handler.UpdateUser)
            users.DELETE("/:id", handler.DeleteUser)
        }
    }
    return router
}
