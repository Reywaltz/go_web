package queries

import (
	"context"
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/internal/models/journal"
)

const (
	selectJoinedJournalQuery = `SELECT t1.id, t1.short_name, t1.in_time, t1.count, t1.mark_name, t1.student_id, t1.surname, t1.name, t1.second_name, t1.group_name FROM (select journal.id, (select short_name from
		subject where journal.study_plan_id = subject.id), journal.in_time, journal.count, (select name from mark where id = journal.mark_id) as mark_name, journal.student_id, student.surname, student.name, student.second_name, (select name from study_group where
		id = student.study_group_id) as group_name from journal inner join student on
		student.id = journal.student_id) as t1`

	selectJoinedJournalByGroupQuery = `SELECT t1.id, t1.short_name, t1.in_time, t1.count, t1.mark_name, t1.student_id, t1.surname, t1.name, t1.second_name, t1.group_name FROM (select journal.id, (select short_name from
		subject where journal.study_plan_id = subject.id), journal.in_time, journal.count, (select name from mark where id = journal.mark_id) as mark_name, journal.student_id, student.surname, student.name, student.second_name, (select name from study_group where
		id = student.study_group_id) as group_name from journal inner join student on
		student.id = journal.student_id) as t1 WHERE t1.group_name=$1`

	selectJoinedJournalByIDQuery = `SELECT t1.id, t1.short_name, t1.in_time, t1.count, t1.mark_name, t1.student_id, t1.surname, t1.name, t1.second_name, t1.group_name FROM (select journal.id, (select short_name from
		subject where journal.study_plan_id = subject.id), journal.in_time, journal.count, (select name from mark where id = journal.mark_id) as mark_name, journal.student_id, student.surname, student.name, student.second_name, (select name from study_group where
		id = student.study_group_id) as group_name from journal inner join student on
		student.id = journal.student_id) as t1 WHERE t1.student_id=$1`

	updateJournalMarkQuery = `update journal set mark_id = $1 where student_id = $2`
)

func (q *Query) Journal() ([]journal.JournalJoined, error) {
	res, err := q.db.Pool().Query(context.Background(), selectJoinedJournalQuery)
	if err != nil {
		log.Fatal("GetALL fatal", err)
		return nil, fmt.Errorf("%w: no Journal data", err)
	}

	out := make([]journal.JournalJoined, 0)
	for res.Next() {
		var journal journal.JournalJoined
		err := res.Scan(&journal.ID,
			&journal.SubjectShortname,
			&journal.InTime,
			&journal.Count,
			&journal.MarkName,
			&journal.StudentID,
			&journal.Surname,
			&journal.Name,
			&journal.SecondName,
			&journal.GroupName)
		if err != nil {
			return nil, err
		}
		log.Println(journal)
		out = append(out, journal)
	}
	return out, nil
}

func (q *Query) GetRecordByGroup(groupName string) ([]journal.JournalJoined, error) {
	res, err := q.db.Pool().Query(context.Background(), selectJoinedJournalByGroupQuery, groupName)
	if err != nil {
		log.Fatal("GetbyGroupALL fatal", err)
		return nil, fmt.Errorf("%w: no Journal data", err)
	}
	out := make([]journal.JournalJoined, 0)
	for res.Next() {
		var journal journal.JournalJoined
		err := res.Scan(&journal.ID,
			&journal.SubjectShortname,
			&journal.InTime,
			&journal.Count,
			&journal.MarkName,
			&journal.StudentID,
			&journal.Surname,
			&journal.Name,
			&journal.SecondName,
			&journal.GroupName)
		if err != nil {
			return nil, err
		}
		out = append(out, journal)
	}
	return out, nil
}

func (q *Query) GetRecordByID(id int) ([]journal.JournalJoined, error) {
	res, err := q.db.Pool().Query(context.Background(), selectJoinedJournalByIDQuery, id)
	if err != nil {
		log.Fatal("GetbyGroupALL fatal", err)
		return nil, fmt.Errorf("%w: no Journal data", err)
	}
	out := make([]journal.JournalJoined, 0)
	for res.Next() {
		var journal journal.JournalJoined
		err := res.Scan(&journal.ID,
			&journal.SubjectShortname,
			&journal.InTime,
			&journal.Count,
			&journal.MarkName,
			&journal.StudentID,
			&journal.Surname,
			&journal.Name,
			&journal.SecondName,
			&journal.GroupName)
		if err != nil {
			return nil, err
		}
		out = append(out, journal)
	}
	return out, nil
}

func (q *Query) UpdateRecord(newJournal journal.Journal) error {
	log.Println(newJournal.MarkID, newJournal.ID)
	_, err := q.db.Pool().Exec(context.Background(), updateJournalMarkQuery, newJournal.MarkID, newJournal.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
