package main

import (
	"fmt"
	"net/http"
)

type counter int

func (c *counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*c++
	fmt.Fprintf(w, "counter = %d\n", *c)
}

func main() {
	var c counter
	c = 0

	http.ListenAndServe("127.0.0.1:1338", &c)
}
