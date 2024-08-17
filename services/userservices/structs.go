package userservices

type UserService struct {
	ServiceName        string `json:"ServiceName"`
	ServiceDescription string `json:"ServiceDescription"`
	ServiceIcon        string `json:"ServiceIconUrl"`
}
