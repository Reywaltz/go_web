package journal

type Journal struct {
	ID          int `json:"id"`
	StudentID   int `json:"student_id"`
	StudyPlanID int `json:"study_plan_id"`
	InTime      int `json:"in_time"`
	Count       int `json:"count"`
	MarkID      int `json:"mark_id"`
}

type JournalJoined struct {
	ID               int    `json:"id"`
	SubjectShortname string `json:"short_name"`
	InTime           int    `json:"in_time"`
	Count            int    `json:"count"`
	MarkName         string `json:"mark_name"`
	StudentID        int    `json:"student_id"`
	Surname          string `json:"surname"`
	Name             string `json:"name"`
	SecondName       string `json:"second_name"`
	GroupName        string `json:"group_name"`
}
