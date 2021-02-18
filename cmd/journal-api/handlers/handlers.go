package handlers

import (
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/internal/studygroup"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll() ([]studygroup.StudyGroup, error)
	Create(studygroup.StudyGroup) error
}

type StudyGroupHandlers struct {
	studyGroupStorage Repository
}

func NewUserGroupHandler(studygroupStorage Repository) *StudyGroupHandlers {
	return &StudyGroupHandlers{
		studyGroupStorage: studygroupStorage,
	}
}

func (h *StudyGroupHandlers) Route(eng *gin.Engine) {

	v1 := eng.Group("/studentgroup")
	{
		v1.GET("", h.getAll)
		v1.POST("", h.createGroup)
	}
}

func (h *StudyGroupHandlers) getAll(c *gin.Context) {
	out, err := h.studyGroupStorage.GetAll()
	if err != nil {
		fmt.Printf("%s can't get studygroup data in handler", err)
		return
	}
	c.JSON(200, out)
}

func (h *StudyGroupHandlers) createGroup(c *gin.Context) {
	var newGroup studygroup.StudyGroup
	c.Bind(&newGroup)

	if newGroup.Name == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "json format error"})
		return
	}
	err := h.studyGroupStorage.Create(newGroup)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"success": "created"})
}
