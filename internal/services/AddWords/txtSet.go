package addwords

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	addwordspg "github.com/antalkon/English_Words_Game/internal/database/addWords_pg"
	"github.com/antalkon/English_Words_Game/internal/models"
	"github.com/antalkon/English_Words_Game/pkg/UUID"
	"github.com/gin-gonic/gin"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func TxtSet(c *gin.Context) {
	SetID, err := UUID.GenerateWordID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createWordSet models.AddSetTxt
	if err := c.Bind(&createWordSet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createWordSet.SetId = SetID
	createWordSet.Code = generateRandomString(16)
	joinDB, err := addwordspg.CreateWordSet(createWordSet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = joinDB

	fileHeader, err := c.FormFile("txt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file", "details": err.Error()})
		return
	}

	// Открываем файл
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file", "details": err.Error()})
		return
	}
	defer file.Close()

	// Создаем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Переменная для хранения строк и подсчета количества строк
	var lines []string
	lineCount := 0

	// Читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineCount++
		fmt.Println(line) // Обработка каждой строки в первом цикле
	}

	// Проверка на наличие ошибок при сканировании
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Второй цикл для дополнительной обработки каждой строки
	for i := 0; i < lineCount; i++ {
		UUID, err := UUID.GenerateWordID()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		line := strings.ReplaceAll(lines[i], " ", "")
		split := strings.Split(line, "-")
		if len(split) == 2 {
			// Обработка каждой строки
			joinDB, err := addwordspg.AddWordSet(createWordSet.Code, UUID, split[0], split[1])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			fmt.Println("Добавлено в базу данных:", joinDB)
		} else {
			log.Printf("Неправильный формат строки: %s\n", line)
		}
	}

	// Возвращаем количество строк в ответе
	c.JSON(http.StatusOK, gin.H{"message": lineCount})
}
