package repository

import (
	"github.com/Reywaltz/web_test/internal/models/journal"
	"github.com/Reywaltz/web_test/internal/models/student"
	"github.com/Reywaltz/web_test/internal/models/studygroup"
)

type StudyGroupRepository interface {
	GetAll() ([]studygroup.StudyGroup, error)
	GetOne(groupName string) (studygroup.StudyGroup, error)
	Create(studygroup.StudyGroup) error
	Delete(groupname string) error
	Update(studyGroup studygroup.StudyGroup) error
	GetGroupByID(groupID int) (studygroup.StudyGroup, error)
}

type StudentRepository interface {
	Students() ([]student.Student, error)
	GetStudentByID(id int) (student.Student, error)
	GetStudentsByGroup(groupName string) ([]student.StudentJoined, error)
	CreateStudent(student student.Student) error
	DeleteStudent(id int) error
	UpdateStudent(student.Student) error
}

type JournalRepository interface {
	Journal() ([]journal.JournalJoined, error)
	GetRecordByGroup(groupName string) ([]journal.JournalJoined, error)
	GetRecordByID(id int) ([]journal.JournalJoined, error)
	UpdateRecord(newJournal journal.Journal) error
}
