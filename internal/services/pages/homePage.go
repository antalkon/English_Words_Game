package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageLoad(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Зентас English",
		"favicon": "static/img/favicon.ico",
	})
}
