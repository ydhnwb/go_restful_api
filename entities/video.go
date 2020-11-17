package entities

// Person struct is used for Video author
type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" binding:"required,email"`
}

// Video is a data class
type Video struct {
	Title       string `json:"title" binding:"min=1,max=100"`
	Description string `json:"description" binding:"max=255"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
