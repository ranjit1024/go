package main

import (
	"fmt"
	"time"
)

func time_formatiing() {
	p := fmt.Println
	t := time.Now()

	p(t)
	p(t.Format(time.Stamp))
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)
	fmt.Println("Hello World...")

}
