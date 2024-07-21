package userdata

import (
	"fmt"
	"net/http"

	datapg "github.com/antalkon/English_Words_Game/internal/database/data_pg"
	jwt "github.com/antalkon/English_Words_Game/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func UserData(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		fmt.Println("err1")
		c.Redirect(http.StatusMovedPermanently, "")
		return
	}
	decode, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		fmt.Println("err2")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	data, err := datapg.UserData(decode.UserID)
	if err != nil {
		fmt.Println("err3")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if data == nil {
		fmt.Println("err4")
		c.JSON(http.StatusNotFound, gin.H{"error": "User data not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"phone":   data.Phone,
		"email":   data.Email,
		"name":    data.Name,
		"surname": data.Surname,
	})
}
