package handlers

import (
	"fmt"

	"github.com/Reywaltz/web_test/internal/studygroup"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll() ([]studygroup.StudyGroup, error)
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
		v1.POST("/qwerty", h.getAll)
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
