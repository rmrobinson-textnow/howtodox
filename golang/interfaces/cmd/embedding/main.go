package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type lockableType struct {
	sync.Mutex

	fieldA map[int]int
	count  int
}

func process(l *lockableType) {
	// Since we have embedded a mutex into lockableType we can use it's methods directly.

	// Try commenting out the lock/unlock method calls and see what happens!
	l.Lock()
	defer l.Unlock()

	idx := rand.Int()
	val := rand.Int()

	fmt.Printf("Round %d, inserting %d to idx %d\n", l.count, val, idx)

	l.fieldA[idx] = val
	l.count++
}

func main() {
	rand.Seed(1)

	l := &lockableType{
		fieldA: map[int]int{},
		count:  0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			process(l)
			wg.Done()
		}()
	}

	wg.Wait()

	for idx, val := range l.fieldA {
		// This print syntax provides a left-pad to 20 characters.
		// Matches the fairly typical C/C++ format style.
		fmt.Printf("%20d = %d\n", idx, val)
	}
}
