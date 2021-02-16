package queries

import (
	"context"

	"github.com/Reywaltz/web_test/internal/studygroup"
	"github.com/Reywaltz/web_test/pkg/postgres"
)

const (
	updateStudyGroupFields = `name`
	selectStudyGroupFields = `id, ` + updateStudyGroupFields
	selectStudyGroupQuery  = `SELECT ` + selectStudyGroupFields + ` FROM study_group`
)

func AllGroups(db *postgres.DB) []studygroup.StudyGroup {
	res, err := db.Session.Query(context.Background(),
		selectStudyGroupQuery)
	if err != nil {
		return nil
	}

	out := make([]studygroup.StudyGroup, 0)
	for res.Next() {
		var studentGroup studygroup.StudyGroup
		err := res.Scan(&studentGroup.ID,
			&studentGroup.Name)
		if err != nil {
			return nil
		}
		out = append(out, studentGroup)
	}
	return out
}
