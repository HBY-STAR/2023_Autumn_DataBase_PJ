package platform

type CreatePlatformRequest struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Country string `json:"country"`
}

// 用platform替代
type UpdatePlatformRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Country string `json:"country"`
}
