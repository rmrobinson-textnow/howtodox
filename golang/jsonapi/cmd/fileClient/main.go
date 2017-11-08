package main

import (
	"fmt"
	"os"

	"github.com/rmrobinson-textnow/howtodox/golang/jsonapi"
	"flag"
	"encoding/json"
)

var (
	filePath = flag.String("filePath", "", "The path to the file to parse")
)

func main() {
	flag.Parse()

	if len(*filePath) < 1 {
		fmt.Printf("Error: supplied file name must exist")
		os.Exit(1)
	}

	// File operations are pretty similar to the standard POSIX methods.
	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Printf("Unable to get open posts file: %s\n", err.Error())
		os.Exit(1)
	}

	// We opened the file above, however we want to make sure it is always closed when we are done.
	// The defer keyword lets us schedule a method call for execution at the end of the execution of the current scope.
	// This guarantees that the file will always be closed.
	defer file.Close()

	// Here we reference the Post defined in the jsonapi package.
	// Since we are in a different directory than where Post is, we need to reference this via it's package.
	var posts []jsonapi.Post

	// The file, that was opened above, is another implementer of the io.Reader interface
	// This allows us to pass it into the json.NewDecoder call without changing the approach.
	err = json.NewDecoder(file).Decode(&posts)

	// If the JSON is malformed for some reason, we catch the parsing errors here.
	if err != nil {
		fmt.Printf("Unable to decode file contents: %s\n", err.Error())
		os.Exit(1)
	}

	// We iterate over the posts and print them to the console.
	for _, post := range posts {
		fmt.Printf("ID: %d, User ID: %d, Title: %s\n", post.ID, post.UserID, post.Title)
	}
}
