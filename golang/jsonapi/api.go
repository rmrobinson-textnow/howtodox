package jsonapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const jsonContentType = "application/json; charset=utf-8"

// API is a handle to an instance of the API.
type API struct {
	url string
}

// NewAPI creates a new instance of an API with the specified URL.
func NewAPI(url string) *API {
	return &API{
		url: url,
	}
}

// GetPosts retrieves a collection of post entries from the API and returns them, or an error.
func (a *API) GetPosts() ([]Post, error) {
	var ret []Post

	// Here we make the HTTP request to the endpoint.
	// The Get method returns 2 results, the Response, and an error.
	// If the error is not nil (i.e. it is set to a value) this function is assumed to have had an issue
	res, err := http.Get(a.url + "/posts")

	if err != nil {
		return ret, err
	}

	// Here we use the built-in JSON decoder type to parse the results.
	// It understands that ret is an array of Posts, and so it simply deserializes the content into this variable.
	// The JSON decoder takes in an io.Reader, which allows us to pass the result body directly in.
	err = json.NewDecoder(res.Body).Decode(&ret)

	return ret, err
}

func (a *API) CreatePost(title string, body string, userID int) (Post, error) {
	ret := Post{}
	newPost := Post{
		UserID: userID,
		Title:  title,
		Body:   body,
	}

	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(newPost)

	if err != nil {
		return ret, err
	}

	res, err := http.Post(a.url+"/posts", jsonContentType, buf)

	if err != nil {
		return ret, err
	}

	err = json.NewDecoder(res.Body).Decode(&ret)

	return ret, err
}
