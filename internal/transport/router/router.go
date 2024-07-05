package router

import (
	"github.com/antalkon/English_Words_Game/internal/transport/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.GET("/", h.HomePage)

	addWordsApi := router.Group("/newWords")
	{
		addWordsApi.POST("/classic", h.ClassicAddWords)
		addWordsApi.POST("/file/txt", h.LoadTxtFile)
	}

	return router

}
