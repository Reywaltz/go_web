package queries

import (
	"context"
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/internal/studygroup"
)

const (
	updateStudyGroupFields = `name`
	selectStudyGroupFields = `id, ` + updateStudyGroupFields
	selectStudyGroupQuery  = `SELECT ` + selectStudyGroupFields + ` FROM study_group`
)

func (q *Query) GetAll() ([]studygroup.StudyGroup, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudyGroupQuery)
	if err != nil {
		log.Fatal("GetALL fatal", err)
		return nil, fmt.Errorf("%w: no group data", err)
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
