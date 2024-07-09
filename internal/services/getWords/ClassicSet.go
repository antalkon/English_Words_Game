package getwords

import (
	"net/http"

	getwordspg "github.com/antalkon/English_Words_Game/internal/database/getWords_pg"
	"github.com/gin-gonic/gin"
)

func ClassicSet(c *gin.Context) {
	set := c.Param("test")
	if set == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'test' is required"})
		return
	}

	word, err := getwordspg.GetRandomSet(set)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"word":      word.Word,
		"translate": word.Translate,
	})
}
