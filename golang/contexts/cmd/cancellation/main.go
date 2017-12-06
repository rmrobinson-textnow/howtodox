package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func simulateNetworkRequestWithContext(ctx context.Context, sleep time.Duration) bool {
	tc := time.NewTicker(sleep)

	for {
		select {
		case <-tc.C:
			fmt.Printf("Request complete\n")
			return true
		case <-ctx.Done():
			fmt.Printf("Context expired, aborting request\n")
			return false
		}
	}
}

// This is an example of a function that would sit in the middle
// of a call flow. It is given a context and can do things
// on it, including invoking other functions that take a context.
// It is coded to take about 3 seconds to complete. you can imagine
// the calls to simulateNetworkRequestWithContext are a couple of outbound
// HTTP requests made to other APIs.
func do3SecondsOfWorkWithAContext(ctx context.Context, data string) {
	fmt.Printf("We're going and doing some work here... work work work\n")

	fmt.Printf("We're working with '%s'\n", data)
	fmt.Printf("We're pretending we're making an outbound request\n")

	_ = simulateNetworkRequestWithContext(ctx, time.Second)

	fmt.Printf("Great! we've gotten some data back and doing some work.\n")

	fmt.Printf("We're going and making another outbound request, this one will be slower\n")

	ok := simulateNetworkRequestWithContext(ctx, time.Second*2)

	if ok {
		fmt.Printf("Great! we got our data back and we're done!\n")
	} else {
		fmt.Printf("Sad face! We didn't complete the request in time\n")
	}
}

func main() {
	// basic use of background context - nothing different happens here
	fmt.Printf("Sleeping with the background context\n")
	fmt.Printf(" Start time: %s\n", time.Now().Format(time.RFC3339))
	simulateNetworkRequestWithContext(context.Background(), time.Second*5)
	fmt.Printf(" End   time: %s\n\n", time.Now().Format(time.RFC3339))

	// more advanced use case - context with a timeout. this also returns a cancel func that I'm ignoring
	fmt.Printf("Sleeping with timeout\n")
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*2)

	fmt.Printf(" Start time: %s\n", time.Now().Format(time.RFC3339))
	simulateNetworkRequestWithContext(timeoutCtx, time.Second*5)
	fmt.Printf(" End   time: %s\n\n", time.Now().Format(time.RFC3339))

	// this allows us to control when the children are cancelled by invoking the returned function.
	fmt.Printf("Sleeping with cancel\n")
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	// this goroutine simulates other work that happens in the background then cancels the context.
	go func() {
		time.Sleep(time.Second)
		fmt.Printf("Cancelling context\n")
		cancelFunc()
	}()

	fmt.Printf(" Start time: %s\n", time.Now().Format(time.RFC3339))
	simulateNetworkRequestWithContext(cancelCtx, time.Second*5)
	fmt.Printf(" End   time: %s\n\n", time.Now().Format(time.RFC3339))

	// here we're simulating more work with a timeout.
	// we expect this one to get cancelled.
	fmt.Printf("Simulating a longer workflow with a 2-second timeout\n")
	workTimeout, _ := context.WithTimeout(context.Background(), time.Second*2)

	fmt.Printf(" Start time: %s\n", time.Now().Format(time.RFC3339))
	do3SecondsOfWorkWithAContext(workTimeout, "here's some work")
	fmt.Printf(" End   time: %s\n\n", time.Now().Format(time.RFC3339))

	// here we're simulating more work with a timeout.
	// we expect this one to complete due to the longer timeout.
	fmt.Printf("Simulating a longer workflow with a 5-second timeout\n")
	workTimeout2, _ := context.WithTimeout(context.Background(), time.Second*5)

	fmt.Printf(" Start time: %s\n", time.Now().Format(time.RFC3339))
	do3SecondsOfWorkWithAContext(workTimeout2, "here's some work 2")
	fmt.Printf(" End   time: %s\n\n", time.Now().Format(time.RFC3339))
}
