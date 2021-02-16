package studyplan

type Studyplan struct {
	ID         int `json:"id"`
	SubjectID  int `json:"subject_id"`
	ExamTypeID int `json:"exam_type_id"`
}
