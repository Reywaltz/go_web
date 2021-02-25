package queries

import (
	"context"
	"fmt"

	"github.com/Reywaltz/web_test/internal/models/student"
)

const (
	joinStudentFields = `student.id, student.surname,
			student.name,student.second_name, study_group.name`

	updateStudentFields = `surname, name, second_name, study_group_id`
	selectStudentFields = `id, ` + updateStudentFields

	selectStudentQuery = `SELECT ` + selectStudentFields +
		` FROM student`

	selectStudentsDebtsQuery = `SELECT t1.id, t1.surname, t1.name, t1.second_name, t1.group as study_group, COUNT(t1.mark_id) FROM (SELECT student.id, student.surname, 
	student.name, student.second_name, study_group.name as group, journal.mark_id
	FROM student INNER JOIN journal ON journal.student_id = student.id INNER JOIN study_group ON study_group.id = student.study_group_id
	WHERE journal.mark_id=4 OR journal.mark_id=7 OR journal.mark_id=6) as t1
	GROUP BY t1.id, t1.surname, t1.name, t1.second_name, t1.group
   `

	selectStudentsMarksQuery = `SELECT DISTINCT t3.student_id as id,
	student.surname, student.name, student.second_name, t3.short_name,
	t3.value as mark_num, t3.name as mark_full, study_group.name as study_group FROM student
	INNER JOIN journal ON student.id = journal.student_id INNER JOIN
	(SELECT mark.value, t2.short_name, t2.student_id, mark.name FROM mark INNER JOIN (SELECT subject.short_name, t1.student_id, t1.mark_id FROM subject INNER JOIN
	(SELECT study_plan.subject_id, journal.student_id, journal.mark_id FROM study_plan INNER JOIN journal
	ON study_plan.id = journal.study_plan_id) as t1 ON subject.id = t1.subject_id) as t2 ON t2.mark_id = mark.id) as t3
	ON t3.student_id = student.id
	INNER JOIN study_group ON student.study_group_id = study_group.id
   `

	selectStudentByGroup = `SELECT ` + joinStudentFields + ` from student 
		INNER JOIN study_group ON study_group.id = student.study_group_id 
		where study_group.name=$1`

	selectStudentByIDQuery = `SELECT ` + selectStudentFields +
		` FROM student WHERE id = $1`

	createStudentQuery = `INSERT INTO student (` + updateStudentFields + `)
							 VALUES ($1, $2, $3, $4)`

	deleteStudentQuery = `DELETE FROM student where id = $1`

	updateStudentQuery = `UPDATE student SET surname=$1, name=$2, second_name=$3, study_group_id=$4 WHERE id = $5`
)

func (q *Query) Students() ([]student.Student, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudentQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: Error in Students", err)
	}

	out := make([]student.Student, 0)
	for res.Next() {
		var student student.Student
		err = res.Scan(&student.ID,
			&student.Surname, &student.Name,
			&student.SecondName, &student.StudyGroupID)
		if err != nil {
			return nil, err
		}

		out = append(out, student)
	}

	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%w: unexpected error", err)
	}

	return out, nil
}

func (q *Query) GetStudentByID(id int) (student.Student, error) {
	var student student.Student

	err := q.db.Pool().QueryRow(
		context.Background(),
		selectStudentByIDQuery, id).Scan(
		&student.ID,
		&student.Surname,
		&student.Name,
		&student.SecondName,
		&student.StudyGroupID,
	)
	if err != nil {
		return student, fmt.Errorf("%w: no group data", err)
	}

	return student, nil
}

func (q *Query) GetStudentsByGroup(groupName string) ([]student.StudentJoined, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudentByGroup, groupName)
	if err != nil {
		return nil, fmt.Errorf("%w: no student data", err)
	}

	out := make([]student.StudentJoined, 0)
	for res.Next() {
		var student student.StudentJoined
		err := res.Scan(&student.ID,
			&student.Surname, &student.Name,
			&student.SecondName, &student.StudyGroup)
		if err != nil {
			return nil, err
		}

		out = append(out, student)
	}

	return out, nil
}

func (q *Query) CreateStudent(student student.Student) error {
	_, err := q.db.Pool().Exec(context.Background(), createStudentQuery,
		student.Surname, student.Name, student.SecondName, student.StudyGroupID)
	if err != nil {
		return fmt.Errorf("%w: error in create", err)
	}

	return nil
}

func (q *Query) DeleteStudent(id int) error {
	_, err := q.db.Pool().Exec(context.Background(), deleteStudentQuery, id)
	if err != nil {
		return fmt.Errorf("%w: error in delete", err)
	}

	return nil
}

func (q *Query) UpdateStudent(student student.Student) error {
	_, err := q.db.Pool().Exec(context.Background(), updateStudentQuery,
		student.Surname,
		student.Name,
		student.SecondName,
		student.StudyGroupID,
		student.ID)
	if err != nil {
		return fmt.Errorf("%w: error in update", err)
	}

	return nil
}

func (q *Query) GetStudentsDebts() ([]student.StudentWithDebts, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudentsDebtsQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: error in select", err)
	}
	out := make([]student.StudentWithDebts, 0)
	for res.Next() {
		var student student.StudentWithDebts

		err = res.Scan(&student.ID, &student.Surname, &student.Name,
			&student.SecondName, &student.StudyGroup, &student.DebtsCount)
		if err != nil {
			return nil, fmt.Errorf("%w: error in scan", err)
		}

		out = append(out, student)
	}
	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%w: unexpected error", err)
	}

	return out, nil
}

func (q *Query) GetStudentMarks() ([]student.StudentwithMarks, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudentsMarksQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: error in select", err)
	}

	out := make([]student.StudentwithMarks, 0)
	for res.Next() {
		var student student.StudentwithMarks

		err = res.Scan(&student.ID,
			&student.Surname,
			&student.Name,
			&student.SecondName,
			&student.ShortName,
			&student.MarkNum,
			&student.MarkFull,
			&student.StudyGroup)
		if err != nil {
			return nil, fmt.Errorf("%w: error in scan", err)
		}

		out = append(out, student)
	}
	if err = res.Err(); err != nil {
		return nil, fmt.Errorf("%w: unexpected error", err)
	}

	return out, nil
}
