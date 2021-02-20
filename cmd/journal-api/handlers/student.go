package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Reywaltz/web_test/internal/repository"
	"github.com/Reywaltz/web_test/internal/student"
	"github.com/gin-gonic/gin"
)

type StudentHandlers struct {
	StudentStorage repository.StudentRepository
}

func NewStudentHandler(studentStorage repository.StudentRepository) *StudentHandlers {
	return &StudentHandlers{
		StudentStorage: studentStorage,
	}
}

func (h *StudentHandlers) Route(eng *gin.Engine) {

	v1 := eng.Group("/student")
	{
		v1.GET("group/:groupName", h.getByGroup)
		v1.GET("", h.getAll)
		v1.GET("id:id", h.getOne)
		v1.POST("", h.Create)
		v1.DELETE("id:id", h.Delete)
		v1.PUT("id:id", h.Update)
	}
}

func (h *StudentHandlers) getAll(c *gin.Context) {

	out, err := h.StudentStorage.Students()
	if err != nil {
		fmt.Printf("%s can't get studygroup data in handler", err)
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
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}
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

func (h *StudentHandlers) Create(c *gin.Context) {
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

func (h *StudentHandlers) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	err = h.StudentStorage.DeleteStudent(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success"})
}

func (h *StudentHandlers) Update(c *gin.Context) {
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
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "success"})
}
