package dto

type CreateResponse struct {
	ShortURL string `json:"short_url"`
}

type GetResponse struct {
	OriginalURL string `json:"original_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
