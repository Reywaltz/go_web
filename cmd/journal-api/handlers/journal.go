package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Reywaltz/web_test/internal/models/journal"
	"github.com/gin-gonic/gin"
)

type JournalRepository interface {
	Journal() ([]journal.JournalJoined, error)
	GetRecordByGroup(groupName string) ([]journal.JournalJoined, error)
	GetRecordByID(id int) ([]journal.JournalJoined, error)
	UpdateRecord(newJournal journal.Journal) error
}

type JournalHandlers struct {
	JournalStorage JournalRepository
}

func NewJournalHandler(journalStorage JournalRepository) *JournalHandlers {
	return &JournalHandlers{
		JournalStorage: journalStorage,
	}
}

func (h *JournalHandlers) Route(eng *gin.Engine) {
	v1 := eng.Group("/journal")
	{
		v1.GET("", h.getAll)
		v1.GET("main/", h.renderHTML)
		v1.GET("group/:groupName", h.getByGroup)
		v1.GET("student/:id", h.getByStudentID)
		v1.PUT("student/:id", h.updateMark)
	}
}

func (h *JournalHandlers) getAll(c *gin.Context) {
	res, err := h.JournalStorage.Journal()
	if err != nil {
		log.Println("Error Journall", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})

		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *JournalHandlers) getByGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	res, err := h.JournalStorage.GetRecordByGroup(groupName)
	if err != nil {
		log.Println("Error JournbyGroup", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})

		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *JournalHandlers) getByStudentID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}
	res, err := h.JournalStorage.GetRecordByID(id)
	if err != nil {
		log.Println("Error JournbyGroup", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, res)
}

func (h *JournalHandlers) updateMark(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}
	var newJournal journal.Journal

	newJournal.ID = id

	c.Bind(&newJournal)
	if newJournal.MarkID == 0 || newJournal.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "wrong json format"})

		return
	}
	err = h.JournalStorage.UpdateRecord(newJournal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "updated"})
}

func (h *JournalHandlers) renderHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}
