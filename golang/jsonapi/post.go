package jsonapi

// Post represents a single post
type Post struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"Title"`
	// Since the type in the message has the same name as this, we don't define a tag explicitly
	Body string
}
