package addwords

import (
	"net/http"

	addwordspg "github.com/antalkon/English_Words_Game/internal/database/addWords_pg"
	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/UUID"
	"github.com/gin-gonic/gin"
)

func AddClassicWords(c *gin.Context) {
	UUID, err := UUID.GenerateWordID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var createWord models.AddClassicWords
	if err := c.ShouldBindJSON(&createWord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createWord.WordId = UUID

	joinDB, err := addwordspg.AddClassicWords(createWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = joinDB
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
