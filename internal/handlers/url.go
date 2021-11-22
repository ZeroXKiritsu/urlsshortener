package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ZeroXKiritsu/urlshortener/structs"
)

func (h *Handler) createShortURL(c *gin.Context) {
	var input structs.Requests

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}

	shortURL, err := h.services.ShortURL.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, structs.ShortURLResponse{
		ShortURL: shortURL,
	})
}

func (h *Handler) getOriginal(c *gin.Context) {
	shortURL := c.Param("short_url")
	if shortURL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}

	original, err := h.services.ShortURL.GetOriginal(shortURL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
		return
	} else if original == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "URL does not exist")
		return
	}

	c.JSON(http.StatusOK, structs.OriginalURLResponse{
		Original: original,
	})
}