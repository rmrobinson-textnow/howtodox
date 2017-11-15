package main

import (
	"flag"
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"sync"
)

var (
	reqCount = flag.Int("reqCount", 100, "The number of connections to make")
	serverURL = flag.String("serverURL", "http://127.0.0.1:1337", "The server URL to conect to")
)

type waitResponse struct {
	Status string `json:"status"`
}

// producer takes in a count and writes that many URLs to the returned channel.
func producer(count int) chan string {
	out := make(chan string)

	go func() {
		for i := 0; i < count; i++ {
			out <- *serverURL + "/wait"
		}
	}()

	return out
}

// consumer reads from the urls channel, and after timing the request to the URL it'll write it to the results channel.
func consumer(urls <-chan string, results chan<- time.Duration) {
	for {
		url, ok := <-urls

		// If the channel is closed we consider ourselves to be done
		if !ok {
			return
		}

		// We track when we started the request
		start := time.Now()

		// Here we make the HTTP request to the endpoint.
		// The Get method returns 2 results, the Response, and an error.
		// If the error is not nil (i.e. it is set to a value) this function is assumed to have had an issue
		res, err := http.Get(url)

		// ...and when we finished the request
		end := time.Now()

		// time.Time yields a Duration type when math is performed on them
		runtime := end.Sub(start)

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			var ret waitResponse

			// Here we use the built-in JSON decoder type to parse the results.
			// It understands that ret is an array of Posts, and so it simply deserializes the content into this variable.
			// The JSON decoder takes in an io.Reader, which allows us to pass the result body directly in.
			err = json.NewDecoder(res.Body).Decode(&ret)

			if err != nil || ret.Status != "done" {
				fmt.Printf("Error: %s, Status: %s\n", err.Error(), ret.Status)
			}
		}

		results <- runtime
	}
}

func main() {
	flag.Parse()

	runtimeResults := make(chan time.Duration)
	var wg sync.WaitGroup

	// We create a channel which we can listen on to get the URLs to download and check
	urls := producer(*reqCount)

	// We create 100 consumers which will read from the 'urls' channel and write their results to runtimeResults
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer(urls, runtimeResults)
		}()
	}

	// We create an anonymous method to read from the runtimeResults channel and keep track of each reported runtime then run it.
	// Once we've processed the expected number of results we will calculate the average duration and then print the result.
	wg.Add(1)

	go func() {
		var runtimes []time.Duration
		count := 0

		for rtr := range runtimeResults {
			runtimes = append(runtimes, rtr)
			count++

			if count >= *reqCount {
				break
			}
		}

		// Once we've finished processing we need to close the channel being used by the consumers
		// so they release their holds on the waitgroup.
		close(urls)

		var avg time.Duration
		for _, rt := range runtimes {
			avg += rt
		}

		avg /= time.Duration(*reqCount)

		fmt.Printf("Average request time: %0.2f seconds\n", avg.Seconds())

		wg.Done()
	}()

	wg.Wait()
}