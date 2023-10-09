package tipos

type Post struct {
	ID     int64  `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
