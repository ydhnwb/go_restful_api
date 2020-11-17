package entities

// Video is a data class
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
