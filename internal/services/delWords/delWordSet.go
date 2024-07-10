package delwords

import (
	"net/http"

	delwordspg "github.com/antalkon/English_Words_Game/internal/database/delWords_pg"
	"github.com/gin-gonic/gin"
)

func DelSet(c *gin.Context) {
	set := c.Param("id")
	if set == "" {
		c.JSON(http.StatusNotFound, gin.H{"more": "Set is not found."})
		return
	}

	err := delwordspg.DelSet(set)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"more": "Set is not found."})
	}

	c.JSON(http.StatusOK, gin.H{"message": "suess"})
}
