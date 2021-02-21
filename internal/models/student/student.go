package student

type Student struct {
	ID           int    `json:"id"`
	Surname      string `json:"surname"`
	Name         string `json:"name"`
	SecondName   string `json:"second_name"`
	StudyGroupID int    `json:"study_group_id"`
}

type StudentJoined struct {
	ID         int    `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	StudyGroup string `json:"study_group"`
}
