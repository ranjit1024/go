package main
import (
	"fmt"
	"time"
)

func timer_demo(){
	timer1 := time.NewTimer(2 * time.Second);

	<- timer1.C;
	fmt.Println("Timer 1 fired")
	timer2 := time.NewTimer(time.Second);
	go func(){
		<- timer2.C;
		fmt.Println("Timer 2 fired");
	}();

	time.Sleep(2 * time.Second)
	fmt.Println("Hello World..");
}