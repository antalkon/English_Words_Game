package router

import (
	"github.com/antalkon/English_Words_Game/internal/transport/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.GET("/", h.HomePage)

	return router

}
