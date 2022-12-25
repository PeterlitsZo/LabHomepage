package viewModel

type UserView struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Extra    string `json:"-"`
}
