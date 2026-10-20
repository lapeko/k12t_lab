package model

type Book struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	AuthorIDs []int64 `json:"author_ids"`
}

type BookResponse struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
