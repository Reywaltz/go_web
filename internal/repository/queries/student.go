package queries

import (
	"context"
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/internal/student"
)

const (
	joinStudentFields = `student.id, student.surname,
			student.name,student.second_name, study_group.name`

	updateStudentFields = `surname, name, second_name, study_group_id`
	selectStudentFields = `id, ` + updateStudentFields

	selectStudentQuery = `SELECT ` + selectStudentFields +
		` FROM student`

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
		log.Fatal("GetALL fatal", err)
		return nil, fmt.Errorf("%w: no student data", err)
	}

	out := make([]student.Student, 0)
	for res.Next() {
		var student student.Student
		err := res.Scan(&student.ID,
			&student.Surname, &student.Name,
			&student.SecondName, &student.StudyGroupID)
		if err != nil {
			return nil, err
		}
		out = append(out, student)
	}
	return out, nil
}

func (q *Query) GetStudentByID(id int) (student.Student, error) {
	var student student.Student
	err := q.db.Pool().QueryRow(context.Background(),
		selectStudentByIDQuery, id).Scan(
		&student.ID,
		&student.Surname,
		&student.Name,
		&student.SecondName,
		&student.StudyGroupID)
	if err != nil {
		return student, fmt.Errorf("%w: no group data", err)
	}
	return student, nil
}

func (q *Query) GetStudentsByGroup(groupName string) ([]student.StudentJoined, error) {
	res, err := q.db.Pool().Query(context.Background(), selectStudentByGroup, groupName)
	if err != nil {
		log.Fatal("GetALL fatal", err)
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
		return err
	}
	return nil
}

func (q *Query) DeleteStudent(id int) error {
	_, err := q.db.Pool().Exec(context.Background(), deleteStudentQuery, id)
	if err != nil {
		return fmt.Errorf("%w: no group data", err)
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
		log.Println(err)
		return err
	}
	return nil
}

// func (q *Query) DeleteGroup(groupName string) error {
// 	_, err := q.db.Pool().Exec(context.Background(), deleteStudentQuery, groupName)
// 	if err != nil {
// 		return fmt.Errorf("%w: no group data", err)
// 	}
// 	return nil
// }

// func (q *Query) Update(Student Student.Student) error {
// 	_, err := q.db.Pool().Exec(context.Background(), updateStudentQuery, Student.Name, Student.ID)
// 	if err != nil {
// 		return fmt.Errorf("%w: no group data", err)
// 	}
// 	return nil
// }

// func (q *Query) GetGroupByID(groupID int) (Student.Student, error) {
// 	var student Student.Student
// 	err := q.db.Pool().QueryRow(context.Background(), selectStudentByIDQuery, groupID).Scan(&student.ID, &student.Name)
// 	if err != nil {
// 		return student, fmt.Errorf("%w: no group data", err)
// 	}
// 	return student, err
// }
