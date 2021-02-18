package repository

import "github.com/Reywaltz/web_test/internal/studygroup"

type Repository interface {
	GetAll() ([]studygroup.StudyGroup, error)
	GetOne(groupName string) (studygroup.StudyGroup, error)
	Create(studygroup.StudyGroup) error
	Delete(groupname string) error
	Update(studyGroup studygroup.StudyGroup) error
	GetGroupByID(groupID int) (studygroup.StudyGroup, error)
}
