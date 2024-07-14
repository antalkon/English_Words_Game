package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WordsPageLoad(c *gin.Context) {
	c.HTML(http.StatusOK, "words.html", gin.H{
		"title":   "Слова - тренировка",
		"favicon": "static/img/favicon.ico",
	})
}
