package oauthapi

import (
	"fmt"
	"net/http"

	jwt "github.com/antalkon/English_Words_Game/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func UserInfoRedirect(c *gin.Context) {
	cookie, err := c.Cookie("access_token_app")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "")
		return
	}
	decode, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		fmt.Println("Decode error")
		c.Redirect(http.StatusMovedPermanently, "")
		return
	}

	zentas_cookie, err := c.Cookie("zentasId_refresh")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "")
		return
	}

	_ = decode
	link := fmt.Sprintf(`http://localhost:8080/auth/api/OAuth/v1/auth/data/info?token=%s`, zentas_cookie)
	c.JSON(http.StatusOK, gin.H{
		"sues": link,
	})
}
