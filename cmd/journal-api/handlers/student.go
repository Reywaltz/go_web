package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Reywaltz/web_test/internal/models/student"
	"github.com/gin-gonic/gin"
)

type StudentRepository interface {
	Students() ([]student.Student, error)
	GetStudentByID(id int) (student.Student, error)
	GetStudentsByGroup(groupName string) ([]student.StudentJoined, error)
	CreateStudent(student student.Student) error
	DeleteStudent(id int) error
	UpdateStudent(student.Student) error
	GetStudentsDebts() ([]student.StudentWithDebts, error)
	GetStudentMarks() ([]student.StudentwithMarks, error)
}

type StudentHandlers struct {
	StudentStorage StudentRepository
}

func NewStudentHandler(studentStorage StudentRepository) *StudentHandlers {
	return &StudentHandlers{
		StudentStorage: studentStorage,
	}
}

func (h *StudentHandlers) Route(eng *gin.Engine) {
	v1 := eng.Group("/student")
	{
		v1.GET("group/:groupName", h.getByGroup)
		v1.GET("mark", h.getmarks)
		v1.GET("", h.getAll)
		v1.GET("id/:id", h.getOne)
		v1.GET("debts", h.getDebts)
		v1.POST("", h.create)
		v1.DELETE("id/:id", h.delete)
		v1.PUT("id/:id", h.update)
	}
}

func (h *StudentHandlers) getAll(c *gin.Context) {
	out, err := h.StudentStorage.Students()
	if err != nil {
		log.Println("can't get studygroup data in handler", err)
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *StudentHandlers) getOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	out, err := h.StudentStorage.GetStudentByID(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println(out)
	c.JSON(http.StatusOK, out)
}

func (h *StudentHandlers) getByGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	log.Println(groupName)
	out, err := h.StudentStorage.GetStudentsByGroup(groupName)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "getByGroup"})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *StudentHandlers) create(c *gin.Context) {
	var newStudent student.Student

	c.Bind(&newStudent)

	if newStudent.Name == "" || newStudent.SecondName == "" || newStudent.StudyGroupID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json format error"})
		return
	}

	err := h.StudentStorage.CreateStudent(newStudent)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (h *StudentHandlers) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	err = h.StudentStorage.DeleteStudent(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success"})
}

func (h *StudentHandlers) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	var newStudent student.Student
	newStudent.ID = id

	c.Bind(&newStudent)

	if newStudent.Name == "" || newStudent.SecondName == "" || newStudent.StudyGroupID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json format error"})
		return
	}

	err = h.StudentStorage.UpdateStudent(newStudent)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "success"})
}

func (h *StudentHandlers) getDebts(c *gin.Context) {
	res, err := h.StudentStorage.GetStudentsDebts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *StudentHandlers) getmarks(c *gin.Context) {
	res, err := h.StudentStorage.GetStudentMarks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, res)
}
