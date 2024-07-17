package handler

import (
	addwords "github.com/antalkon/English_Words_Game/internal/services/AddWords"
	"github.com/antalkon/English_Words_Game/internal/services/api"
	delwords "github.com/antalkon/English_Words_Game/internal/services/delWords"
	getwords "github.com/antalkon/English_Words_Game/internal/services/getWords"
	"github.com/antalkon/English_Words_Game/internal/services/pages"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage(c *gin.Context) {
	pages.HomePageLoad(c)
}
func (h *Handler) WordsPage(c *gin.Context) {
	pages.WordsPageLoad(c)
}
func (h *Handler) WordsGamePage(c *gin.Context) {
	pages.WordsPageGame(c)
}

func (h *Handler) ClassicAddWords(c *gin.Context) {
	addwords.AddClassicWords(c)
}

func (h *Handler) LoadTxtFile(c *gin.Context) {
	addwords.LoadTxtFile(c)
}

func (h *Handler) TxtSet(c *gin.Context) {
	addwords.TxtSet(c)
}

func (h *Handler) GetClassicWords(c *gin.Context) {
	getwords.ClassicWords(c)
}

func (h *Handler) GetClassicSet(c *gin.Context) {
	getwords.ClassicSet(c)
}

func (h *Handler) DelSet(c *gin.Context) {
	delwords.DelSet(c)
}

func (h *Handler) ClassicDelWords(c *gin.Context) {
	delwords.DelClassic(c)
}
func (h *Handler) AddPendingEmail(c *gin.Context) {
	api.PnadingEmail(c)
}
func (h *Handler) AddPendingEmail(c *gin.Context) {
	api.PnadingEmail(c)
}
