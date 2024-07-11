package delwords

import (
	"net/http"

	delwordspg "github.com/antalkon/English_Words_Game/internal/database/delWords_pg"
	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/gin-gonic/gin"
)

func DelClassic(c *gin.Context) {

	var delWord models.DelClassicWord
	if err := c.ShouldBindJSON(&delWord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word, err := delwordspg.DelClassicWords(delWord)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"404": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"word":    word,
	})
}
