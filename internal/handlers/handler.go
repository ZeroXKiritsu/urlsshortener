package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ZeroXKiritsu/urlshortener/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler {
	  services: services,
	}
  }
  
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
  
	router.GET("/:short_url", h.getOriginal)
	router.POST("/", h.createShortURL)
  
	return router
}