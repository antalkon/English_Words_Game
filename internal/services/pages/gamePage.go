package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WordsPageGame(c *gin.Context) {
	c.HTML(http.StatusOK, "wordsGame.html", gin.H{
		"title":   "Игра слова",
		"favicon": "static/img/favicon.ico",
	})
}
