package main
import (
"fmt"
"time"
)
func worker(id int, jobs <- chan int, results chan <- int16){
	for j := range jobs {
		fmt.Println("Worker", id , "Started Job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id , "Finished Job")
		results <- j * 2;
	}
}

func demo_worker_pool(){
	fmt.Println("Data")
}