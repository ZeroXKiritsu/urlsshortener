package structs

type Requests struct {
	URL string `json:"url" binding:"required"`
}

type ShortURLResponse struct {
	ShortURL string `json:"short_url"`
}

type OriginalURLResponse struct {
	Original string `json:"original"`
}
