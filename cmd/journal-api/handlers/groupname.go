package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Reywaltz/web_test/internal/models/studygroup"
	"github.com/Reywaltz/web_test/internal/repository"
	"github.com/gin-gonic/gin"
)

type StudyGroupHandlers struct {
	studyGroupStorage repository.StudyGroupRepository
}

func NewStudyGroupHandler(studygroupStorage repository.StudyGroupRepository) *StudyGroupHandlers {
	return &StudyGroupHandlers{
		studyGroupStorage: studygroupStorage,
	}
}

func (h *StudyGroupHandlers) Route(eng *gin.Engine) {

	v1 := eng.Group("/studentgroup")
	{
		v1.GET("", h.getAll)
		v1.GET(":groupName", h.getGroup)
		v1.POST("", h.createGroup)
		v1.DELETE(":groupname", h.deleteGroup)
		v1.PUT(":id", h.updateGroup)
	}
}

func (h *StudyGroupHandlers) getAll(c *gin.Context) {
	out, err := h.studyGroupStorage.GetAll()
	if err != nil {
		fmt.Printf("%s can't get studygroup data in handler", err)
		return
	}
	c.JSON(http.StatusAccepted, out)
}

func (h *StudyGroupHandlers) getGroup(c *gin.Context) {
	groupName := c.Param("groupName")
	log.Println(groupName)
	res, err := h.studyGroupStorage.GetOne(groupName)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, res)
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": "created"})
}

func (h *StudyGroupHandlers) deleteGroup(c *gin.Context) {
	groupName := c.Param("groupname")
	err := h.studyGroupStorage.Delete(groupName)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No such group"})
	}
	c.JSON(200, gin.H{"success": "deleted"})
}

func (h *StudyGroupHandlers) updateGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid group id"})
		return
	}

	tmp, err := h.studyGroupStorage.GetGroupByID(groupID)
	if err != nil {
		log.Println(err)
	}

	if tmp == (studygroup.StudyGroup{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Group not found"})
		return
	}

	newGroup := studygroup.StudyGroup{
		ID:   groupID,
		Name: "",
	}
	c.Bind(&newGroup)

	if newGroup.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Group name not provided"})
		return
	}

	err = h.studyGroupStorage.Update(newGroup)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Smth wrong with db"})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": "updated"})
}
