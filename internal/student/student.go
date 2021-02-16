package student

type Student struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	SecondName   string `json:"second_name"`
	StudyGroupID int    `json:"study_group_id"`
}
