package api

import (
	"net/http"

	apipg "github.com/antalkon/English_Words_Game/internal/database/api_pg"
	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/UUID"
	"github.com/gin-gonic/gin"
)

func PnadingEmail(c *gin.Context) {
	UUID, err := UUID.GenerateWordID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pending models.PendingEmail
	if err := c.ShouldBindJSON(&pending); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pending.Id = UUID
	db, err := apipg.PendingDB(pending)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = db
	c.JSON(http.StatusOK, gin.H{"suesses": "Вы подписаны на рассылку!"})
}
