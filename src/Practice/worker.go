package main
import (
"fmt"
"time"
)
func worker_demo(id int, jobs <- chan int, results chan <- int){
	for j := range jobs {
		fmt.Println("Worker", id , "Started Job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id , "Finished Job")
		results <- j * 2;
	}
}

func demo_worker_pool(){
	const numJobs = 5;
	jobs := make(chan int, numJobs)
	results := make(chan int , numJobs)

	for w := 1; w <= 3; w++{
		go worker_demo(w, jobs, results)
	}

	for j := 

	fmt.Println("Data")
}

Hii Sammmmmu
i just want to tell you that you have a butiful smile,
and when you bit your lips you look cute 
one thing i want to tell for very long time about my felling about you
we stated to wokring in that time i just completey attractes your personalty
and i feel like i just to spend more time with you understand you
or in way i statetd to falling love with you 
i know the felling cannot be mutual,and i know yoo just went to recent break up 
but i know if i dont tell you now i might loost the oppertnite to tell you and 
becase i few months we might end up someware else
i just want to be with you spend time listen to whatever nonsecne you are saying 
and also irrepctive whatever you answer i will make sure our friendship should not 
get hammper...
Happly Birthday
I love you