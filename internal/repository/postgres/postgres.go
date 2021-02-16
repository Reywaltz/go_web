package postgres

import (
	"context"

	"github.com/Reywaltz/web_test/internal/repository"
	"github.com/Reywaltz/web_test/internal/studygroup"
	"github.com/Reywaltz/web_test/pkg/postgres"
)

const (
	updateStudyGroupFields = `name`
	selectStudyGroupFields = `id, ` + updateStudyGroupFields
	selectStudyGroupQuery  = `SELECT ` + selectStudyGroupFields + ` FROM study_group`
)

type Repo struct {
	db *postgres.DB
}

func NewRepository(db *postgres.DB) repository.Repository {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetAll() ([]studygroup.StudyGroup, error) {
	res, err := r.db.Session.Query(context.Background(),
		selectStudyGroupQuery)
	if err != nil {
		return nil, err
	}

	out := make([]studygroup.StudyGroup, 0)
	for res.Next() {
		var studentGroup studygroup.StudyGroup
		err := res.Scan(&studentGroup.ID,
			&studentGroup.Name)
		if err != nil {
			return nil, err
		}
		out = append(out, studentGroup)
	}
	return out, nil
}
