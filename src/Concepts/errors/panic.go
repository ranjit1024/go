package main

import "fmt"

func processJob(jobID int) {
	fmt.Println("Starting job:", jobID)

	if jobID == 3 {
		panic("job failed unexpectedly")
	}

	fmt.Println("Job completed:", jobID)
}

func panic_demo() {
	for i := 1; i <= 5; i++ {
		processJob(i)
	}

	fmt.Println("All jobs finished")
}
