package getwords

import (
	"math/rand"
	"net/http"

	getwordspg "github.com/antalkon/English_Words_Game/internal/database/getWords_pg"
	"github.com/gin-gonic/gin"
)

func ClassicWords(c *gin.Context) {
	// count, err := getwordspg.GetStringCount()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file", "details": err.Error()})
	// }
	// s1 := generateRandomFour(count)
	word, err := getwordspg.GetRandomWord()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file", "details": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"word":      word.Word,
		"translate": word.Translate,
	})

}

func generateRandomFour(count int) int {
	s1 := rand.Intn(count) + 1

	return s1

}
