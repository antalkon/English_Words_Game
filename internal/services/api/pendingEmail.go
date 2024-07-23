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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Не удолось выпустить ID."})
		return
	}
	var pending models.PendingEmail
	if err := c.ShouldBindJSON(&pending); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Не удалось обработать данные."})
		return
	}
	pending.Id = UUID
	db, err := apipg.PendingDB(pending)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Вы уже подписаны."})
		return
	}
	_ = db
	c.JSON(http.StatusOK, gin.H{"message": "Вы подписаны на рассылку!"})
}
