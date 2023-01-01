package viewModel

type ResourceView struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}
