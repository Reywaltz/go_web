package repository

import "github.com/Reywaltz/web_test/internal/studygroup"

type Repository interface {
	GetAll() ([]studygroup.StudyGroup, error)
}
