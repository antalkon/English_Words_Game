package router

import (
	"net/http"

	"github.com/antalkon/English_Words_Game/internal/transport/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("web/public/*/*")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", nil)
	})

	router.Static("/static", "web/static")
	// router.GET("/", h.HomePage)

	pagesWords := router.Group("/")
	{
		pagesWords.GET("/", h.HomePage)
		pagesWords.GET("/words", h.WordsPage)
		pagesWords.GET("/words/game", h.WordsGamePage)

	}
	addWordsApi := router.Group("/newWords")
	{
		addWordsApi.POST("/classic", h.ClassicAddWords)
		addWordsApi.DELETE("/classic", h.ClassicDelWords)
		addWordsApi.POST("/file/txt", h.LoadTxtFile)
		addWordsApi.POST("/set/file/txt", h.TxtSet)
		addWordsApi.DELETE("/set/file/txt/:id", h.DelSet)

	}
	getWordsApi := router.Group("words")
	{
		getWordsApi.POST("/classic", h.GetClassicWords)
		getWordsApi.POST("/classic/set/:test", h.GetClassicSet)
	}

	auth := router.Group("/auth")
	{
		authApi := auth.Group("/api")
		{
			authApi_v1 := authApi.Group("/v1")
			{
				OAuth := authApi_v1.Group("/OAuth")
				{
					OAuth.GET("/link", h.ParseOauthLink)
					OAuth_data := OAuth.Group("/data")
					{
						OAuth_data.GET("/userInfo", h.UserInfoRedirect)
					}
				}
			}
		}
	}
	API := router.Group("api")
	{
		v1 := API.Group("v1")
		{
			v1.POST("/email/pending", h.AddPendingEmail)
			v1.POST("/data/user", h.UserData)
		}
	}
	return router

}
