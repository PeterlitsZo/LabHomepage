package viewModel

type UserView struct {
	ID       uint   `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Extra    string `json:"extra"`
}
