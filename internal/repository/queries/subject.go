package queries

import (
	"context"
	"fmt"

	"github.com/Reywaltz/web_test/internal/models/subject"
)

const (
	getSubjectsAttestationQuery = `SELECT t1.short_name, t1.name, exam_type.type FROM (SELECT subject.name, subject.short_name, study_plan.exam_type_id
		FROM subject INNER JOIN study_plan ON subject.id = study_plan.subject_id) as t1
		INNER JOIN exam_type ON exam_type.id = t1.exam_type_id ORDER BY t1.short_name`
)

func (q *Query) GetSubjectAttestation() ([]subject.Subject, error) {
	res, err := q.db.Pool().Query(context.Background(), getSubjectsAttestationQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: error in select query", err)
	}

	out := make([]subject.Subject, 0)
	for res.Next() {
		var subject subject.Subject
		err = res.Scan(&subject.ShortName, &subject.Name, &subject.Type)
		if err != nil {
			return nil, fmt.Errorf("%w: error during scan", err)
		}

		out = append(out, subject)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%w: unexpected error", err)
	}

	return out, nil
}
