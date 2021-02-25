package handlers

import (
	"log"
	"net/http"

	"github.com/Reywaltz/web_test/internal/models/subject"
	"github.com/gin-gonic/gin"
)

type SubjectRepository interface {
	GetSubjectAttestation() ([]subject.Subject, error)
}

type SubjectHandlers struct {
	SubjectStorage SubjectRepository
}

func NewSubjectHandler(subjectStorage SubjectRepository) *SubjectHandlers {
	return &SubjectHandlers{
		SubjectStorage: subjectStorage,
	}
}

func (h *SubjectHandlers) Route(eng *gin.Engine) {
	v1 := eng.Group("/subject")
	{
		v1.GET("", h.subjects)
	}
}

func (h *SubjectHandlers) subjects(c *gin.Context) {
	out, err := h.SubjectStorage.GetSubjectAttestation()
	if err != nil {
		log.Println("can't get studygroup data in handler", err)

		return
	}
	c.JSON(http.StatusOK, out)
}
