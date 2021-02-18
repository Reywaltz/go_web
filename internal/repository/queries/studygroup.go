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

	selectStudyGroupQuery = `SELECT ` + selectStudyGroupFields +
		` FROM study_group`

	selectOneStudyGroupQuery = `SELECT ` + selectStudyGroupFields +
		` FROM study_group WHERE name = $1`

	selectStudyGroupByIDQuery = `SELECT ` + selectStudyGroupFields +
		` FROM study_group WHERE id = $1`

	createStudyGroupQuery = `INSERT INTO study_group (` + updateStudyGroupFields + `)
							 VALUES ($1)`

	deleteStudyGroupQuery = `DELETE FROM study_group where ` + updateStudyGroupFields + `= $1`

	updateStudyGroupQuery = `UPDATE study_group SET ` + updateStudyGroupFields + `=$1 WHERE id = $2`
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

func (q *Query) GetOne(groupName string) (studygroup.StudyGroup, error) {
	var studentgroup studygroup.StudyGroup
	err := q.db.Pool().QueryRow(context.Background(), selectOneStudyGroupQuery, groupName).Scan(&studentgroup.ID, &studentgroup.Name)
	if err != nil {
		return studentgroup, fmt.Errorf("%w: no group data", err)
	}
	return studentgroup, nil
}

func (q *Query) Create(studygroup studygroup.StudyGroup) error {
	_, err := q.db.Pool().Exec(context.Background(), createStudyGroupQuery, studygroup.Name)
	if err != nil {
		return fmt.Errorf("%w: no group data", err)
	}
	return nil
}

func (q *Query) Delete(groupName string) error {
	_, err := q.db.Pool().Exec(context.Background(), deleteStudyGroupQuery, groupName)
	if err != nil {
		return fmt.Errorf("%w: no group data", err)
	}
	return nil
}

func (q *Query) Update(studyGroup studygroup.StudyGroup) error {
	_, err := q.db.Pool().Exec(context.Background(), updateStudyGroupQuery, studyGroup.Name, studyGroup.ID)
	if err != nil {
		return fmt.Errorf("%w: no group data", err)
	}
	return nil
}

func (q *Query) GetGroupByID(groupID int) (studygroup.StudyGroup, error) {
	var studentgroup studygroup.StudyGroup
	err := q.db.Pool().QueryRow(context.Background(), selectStudyGroupByIDQuery, groupID).Scan(&studentgroup.ID, &studentgroup.Name)
	if err != nil {
		return studentgroup, fmt.Errorf("%w: no group data", err)
	}
	return studentgroup, err
}
