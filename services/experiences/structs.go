package experiences

type Experience struct {
	CompanyName  string `json:"company_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CurrentPlace bool   `json:"current_place"`
	Position     string `json:"position"`
}
