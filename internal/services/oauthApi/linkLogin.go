package oauthapi

import (
	"fmt"
	"net/http"

	oauthpg "github.com/antalkon/English_Words_Game/internal/database/oauth_pg"
	"github.com/antalkon/English_Words_Game/internal/models"
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

	checkDB, err := oauthpg.CheckDBAccByID(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if checkDB == "false" {
		// Преобразуем userData в models.OauthUser
		oauthUser := models.OauthUser{
			ID:      userData.ID,
			Email:   userData.Email,
			Phone:   userData.Phone,
			Name:    userData.Name,
			Surname: userData.Surname,
			Exp:     userData.Exp,
		}

		_, err := oauthpg.AddUser(oauthUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    userData.ID,
		"email": userData.Email,
	})
}
