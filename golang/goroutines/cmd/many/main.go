package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	count = flag.Int("count", 10000, "Number of goroutines to run")
)

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func(id int) {
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Microsecond)
			fmt.Printf("goroutine %d done\n", id)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
