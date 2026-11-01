package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu      sync.Mutex
	countrs map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.countrs[name]++
}

func demo_mutex() {
	c := Container{
		countrs: map[string]int{"a": 1, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 1)
	})

	wg.Wait()
	fmt.Println("Mutexm", c.countrs)

}
