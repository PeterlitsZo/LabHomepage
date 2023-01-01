package viewModel

type PersonView struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}
