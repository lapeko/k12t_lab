package model

type CreateAuthorsRequest struct {
	Names []string `json:"names"`
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
