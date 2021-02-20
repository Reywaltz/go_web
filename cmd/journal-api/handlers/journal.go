package handlers

import (
	"github.com/Reywaltz/web_test/internal/repository"
	"github.com/gin-gonic/gin"
)

type JournalHandlers struct {
	JournalStorage repository.JournalRepository
}

func NewJournalHandler(journalStorage repository.JournalRepository) *JournalHandlers {
	return &JournalHandlers{
		JournalStorage: journalStorage,
	}
}

func (h *JournalHandlers) Route(eng *gin.Engine) {

	v1 := eng.Group("/journal")
	{
		v1.GET("", h.getAll)
		v1.GET("student/:student", h.getAll)
		v1.GET("group/:group", h.getAll)
		v1.PUT("student/:id", h.getAll)
	}
}

func (h *JournalHandlers) getAll(c *gin.Context) {
	c.JSON(200, gin.H{"test": "test"})
}
