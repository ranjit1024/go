package main
import "fmt"

func demo_closing_chan(){
	jobs := make(chan int, 5);
	done := make(chan bool);
	go func(){
		for {
			j , more := <- jobs;
			if more {
				fmt.Println("Recieced Job", j);
			} else{
				fmt.Println("All jobs recived");
				done <- true;
				return;
			}
		}
	}();

	for j := 1; j <= 3; j++{
		jobs <- j;
		fmt.Println("Sent job");
	}
	close(jobs);
	fmt.Println("Send all jobs");
	<- done;
	fmt.Println("Hello World...");
}