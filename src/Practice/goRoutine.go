package main
import ("fmt"
"time"
)

func f(from string){
	for i:=range 3{
		fmt.Println(from, ":", i)
	}
}
func demo_goroutine(){
	go func(msg string){
		fmt.Println(msg)
	}("Hellow World...")
	go f("Goroutine")
	time.Sleep(time.Second)
	fmt.Println("Hellow World")

}