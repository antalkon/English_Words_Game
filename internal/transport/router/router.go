package router

import (
	"github.com/antalkon/English_Words_Game/internal/transport/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()

	// Загрузка основного шаблона base.html и всех остальных шаблонов из папки web/components
	router.LoadHTMLGlob("web/components/*")

	router.GET("/", h.HomePage)
	router.Static("/static", "web/static")

	addWordsApi := router.Group("/newWords")
	{
		addWordsApi.POST("/classic", h.ClassicAddWords)
		addWordsApi.DELETE("/classic", h.ClassicDelWords)
		addWordsApi.POST("/file/txt", h.LoadTxtFile)
		addWordsApi.POST("/set/file/txt", h.TxtSet)
		addWordsApi.DELETE("/set/file/txt/:id", h.DelSet)
	}

	getWordsApi := router.Group("/words")
	{
		getWordsApi.POST("/classic", h.GetClassicWords)
		getWordsApi.POST("/classic/set/:test", h.GetClassicSet)
	}

	return router
}
