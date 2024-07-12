package homepage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageLoad загружает главную страницу с шаблонами header и footer.
func HomePageLoad(c *gin.Context) {
	// Отображение главной страницы с использованием шаблона base.html
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Зентас English",
		"favicon": "static/img/favicon.png",
	})
}
