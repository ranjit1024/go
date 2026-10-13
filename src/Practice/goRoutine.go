package main
import ("fmt"
"time"
)
func demo_goroutine(){
	go func(){
		fmt.Println("data")
	}()
	time.Sleep(time.Second)
	fmt.Println("Hellow World")

}