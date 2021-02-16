package journal

type Journal struct {
	ID          int `json:"id"`
	studentID   int `json:"student_id"`
	studyPlanID int `json:"study_plan_id"`
	inTime      int `json:"in_time`
	count       int `json:"count"`
	markID      int `json:"mark_id"`
}
