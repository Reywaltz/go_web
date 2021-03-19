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

type StudentWithDebts struct {
	ID         int    `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	StudyGroup string `json:"study_group"`
	DebtsCount int    `json:"debts"`
}

type StudentwithMarks struct {
	ID         int    `json:"id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	ShortName  string `json:"short_name"`
	MarkNum    string `json:"mark_num"`
	MarkFull   string `json:"mark_full"`
	StudyGroup string `json:"study_group"`
}

type JoinedFullNameStudent struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullname"`
	StudyGroup string `json:"name"`
}

type UpdateStudentGroup struct {
	StudentID int `json:"student_id"`
	GroupID   int `json:"group_id"`
}
