package educations

type Education struct {
	InstitueName      string `json:"institue_name"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	ModeOfStudy       string `json:"mode_of_study"`
	DegreeType        string `json:"degree_type"`
	AreaOfStudy       string `json:"area_of_study"`
	CurrentlyStudying bool   `json:"currently_studying"`
	Description       string `json:"description"`
}
