package oauthapi

import (
	"fmt"
	"net/http"

	jwt "github.com/antalkon/English_Words_Game/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func ParseLoginLink(c *gin.Context) {
	token := c.Query("q")
	userData, err := jwt.DecodeRefreshToken(token)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    userData.ID,
		"email": userData.Email,
	})
}
