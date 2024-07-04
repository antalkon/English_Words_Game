package handler

import (
	"github.com/antalkon/English_Words_Game/internal/services/pages"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage(c *gin.Context) {
	pages.HomePage(c)
}
