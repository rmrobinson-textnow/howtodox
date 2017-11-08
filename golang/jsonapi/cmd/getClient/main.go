package main

import (
	"fmt"
	"os"

	"github.com/rmrobinson-textnow/howtodox/golang/jsonapi"
)

func main() {
	// We create a new instance of the json API with the specified base URL.
	api := jsonapi.NewAPI("https://jsonplaceholder.typicode.com")

	// We get the posts, and check for an error.
	posts, err := api.GetPosts()

	// If we received an error, print it and exit.
	if err != nil {
		fmt.Printf("Unable to get posts: %s\n", err.Error())
		os.Exit(1)
	}

	// Iterate over the posts and print them.
	for _, post := range posts {
		fmt.Printf("ID: %d, User ID: %d, Title: %s\n", post.ID, post.UserID, post.Title)
	}
}
