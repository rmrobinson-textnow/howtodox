package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rmrobinson-textnow/howtodox/golang/jsonapi"
)

var (
	title = flag.String("title", "", "The title of the post")
	body  = flag.String("body", "", "The body of the post")
)

func main() {
	flag.Parse()

	// We create a new instance of the json API with the specified base URL.
	api := jsonapi.NewAPI("https://jsonplaceholder.typicode.com")

	fmt.Printf("Creating post with title '%s' and body '%s'\n", *title, *body)

	// We get the posts, and check for an error.
	post, err := api.CreatePost(*title, *body, 10)

	// If we received an error, print it and exit.
	if err != nil {
		fmt.Printf("Unable to create post: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("API accepted request, created post ID: %d\n", post.ID)
}
