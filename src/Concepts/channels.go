package main

import "fmt"

func addTen(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n + 10
	}
	close(out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}
func main() {
	fmt.Println("Data is the king ")
	input := make(chan int)
	stage1 := make(chan int)
	stage2 := make(chan int)
	output := make(chan int)

	go addTen(input, stage1)
	go multiplyByTwo(stage1, stage2)
	go square(stage2, output)

	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	for result := range output {
		fmt.Println(result)
	}

}
